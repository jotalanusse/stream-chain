CREATE OR REPLACE FUNCTION klyra_block_processor_unordered_handlers(block jsonb) RETURNS jsonb[] AS $$
/**
  Processes each event that should be handled by the batched handler. This includes all supported non synchronous types

  Parameters:
    - block: A 'DecodedIndexerTendermintBlock' converted to JSON format. Conversion to JSON is expected to be done by JSON.stringify.

  Returns:
    An array containing the results for each event or NULL if this event is not handled by this block processor.
    See each individual handler function for a description of the the inputs and outputs.

  (Note that no text should exist before the function declaration to ensure that exception line numbers are correct.)

  TODO(IND-514): Remove the batch and sync handlers completely by moving all redis updates into
  a pipeline similar to how we return kafka events and then batch and emit them.
*/
DECLARE
    TDAI_ASSET_ID constant text = '0';

    block_height int = (block->'height')::int;
    block_time timestamp = (block->>'time')::timestamp;
    event_ jsonb;
    rval jsonb[];
    event_index int;
    transaction_index int;
    event_data jsonb;
BEGIN
    rval = array_fill(NULL::jsonb, ARRAY[coalesce(jsonb_array_length(block->'events'), 0)]::integer[]);

    /** Note that arrays are 1-indexed in PostgreSQL and empty arrays return NULL for array_length. */
    FOR i in 1..coalesce(array_length(rval, 1), 0) LOOP
        event_ = jsonb_array_element(block->'events', i-1);
        transaction_index = klyra_tendermint_event_to_transaction_index(event_);
        event_index = (event_->'eventIndex')::int;
        event_data = event_->'dataBytes';
        CASE event_->'subtype'
            WHEN '"order_fill"'::jsonb THEN
                /** If event_data.order is populated then this means it is not a liquidation order. */
                IF event_data->'order' IS NOT NULL THEN
                    rval[i] = jsonb_build_object(
                            'makerOrder',
                            klyra_order_fill_handler_per_order('makerOrder', block_height, block_time, event_data, event_index, transaction_index, jsonb_array_element_text(block->'txHashes', transaction_index), 'MAKER', 'LIMIT', TDAI_ASSET_ID, event_data->>'makerCanceledOrderStatus'),
                            'order',
                            klyra_order_fill_handler_per_order('order', block_height, block_time, event_data, event_index, transaction_index, jsonb_array_element_text(block->'txHashes', transaction_index), 'TAKER', 'LIMIT', TDAI_ASSET_ID, event_data->>'takerCanceledOrderStatus'));
                ELSE
                    rval[i] = jsonb_build_object(
                            'makerOrder',
                            klyra_liquidation_fill_handler_per_order('makerOrder', block_height, block_time, event_data, event_index, transaction_index, jsonb_array_element_text(block->'txHashes', transaction_index), 'MAKER', 'LIQUIDATION', TDAI_ASSET_ID),
                            'liquidationOrder',
                            klyra_liquidation_fill_handler_per_order('liquidationOrder', block_height, block_time, event_data, event_index, transaction_index, jsonb_array_element_text(block->'txHashes', transaction_index), 'TAKER', 'LIQUIDATED', TDAI_ASSET_ID));
                END IF;
            WHEN '"subaccount_update"'::jsonb THEN
                rval[i] = klyra_subaccount_update_handler(block_height, block_time, event_data, event_index, transaction_index);
            WHEN '"transfer"'::jsonb THEN
                rval[i] = klyra_transfer_handler(block_height, block_time, event_data, event_index, transaction_index, jsonb_array_element_text(block->'txHashes', transaction_index));
            WHEN '"stateful_order"'::jsonb THEN
                rval[i] = klyra_stateful_order_handler(block_height, block_time, event_data);
            WHEN '"deleveraging"'::jsonb THEN
                rval[i] = klyra_deleveraging_handler(block_height, block_time, event_data, event_index, transaction_index, jsonb_array_element_text(block->'txHashes', transaction_index));
            WHEN '"yield_params"'::jsonb THEN
                rval[i] = klyra_yield_params_handler(block_height, block_time, event_data);
            ELSE
                NULL;
            END CASE;
    END LOOP;

    RETURN rval;
END;
$$ LANGUAGE plpgsql;
