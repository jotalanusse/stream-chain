import {
  OraclePriceColumns,
  OraclePriceCreateObject,
  OraclePriceFromDatabase,
  Ordering,
  PriceMap,
} from '../../src/types';
import * as OraclePriceTable from '../../src/stores/oracle-price-table';
import { clearData, migrate, teardown } from '../../src/helpers/db-helpers';
import { seedData } from '../helpers/mock-generators';
import {
  createdDateTime,
  defaultBlock,
  defaultMarket,
  defaultMarket2,
  defaultOraclePrice,
  defaultOraclePrice2,
  defaultOraclePriceId,
} from '../helpers/constants';
import * as BlockTable from '../../src/stores/block-table';
import { DateTime } from 'luxon';

describe('Oracle price store', () => {
  const updatedHeight: string = '5';

  beforeEach(async () => {
    await seedData();
    await BlockTable.create({
      ...defaultBlock,
      blockHeight: updatedHeight,
    });
  });

  beforeAll(async () => {
    await migrate();
  });

  afterEach(async () => {
    await clearData();
  });

  afterAll(async () => {
    await teardown();
  });

  it('Successfully creates an oracle price', async () => {
    await OraclePriceTable.create(defaultOraclePrice);
  });

  it('Successfully creates multiple oracle prices', async () => {
    const oraclePrice2: OraclePriceCreateObject = {
      marketId: defaultMarket.id,
      spotPrice: '10000.05',
      pnlPrice: '10000.05',
      effectiveAt: createdDateTime.toISO() ?? '',
      effectiveAtHeight: updatedHeight,
    };
    await Promise.all([
      OraclePriceTable.create(defaultOraclePrice),
      OraclePriceTable.create(oraclePrice2),
    ]);

    const oraclePrices: OraclePriceFromDatabase[] = await OraclePriceTable.findAll(
      {
        marketId: [defaultMarket.id],
      },
      [],
      {
        orderBy: [[OraclePriceColumns.effectiveAtHeight, Ordering.ASC]],
      },
    );

    expect(oraclePrices.length).toEqual(2);
    expect(oraclePrices[0]).toEqual(expect.objectContaining(defaultOraclePrice));
    expect(oraclePrices[1]).toEqual(expect.objectContaining(oraclePrice2));
  });

  it('Successfully finds all OraclePrices', async () => {
    await Promise.all([
      OraclePriceTable.create(defaultOraclePrice),
      OraclePriceTable.create({
        ...defaultOraclePrice,
        effectiveAtHeight: updatedHeight,
      }),
    ]);

    const oraclePrices: OraclePriceFromDatabase[] = await OraclePriceTable.findAll(
      {
        marketId: [defaultMarket.id],
      },
      [],
      {
        orderBy: [[OraclePriceColumns.effectiveAtHeight, Ordering.ASC]],
      },
    );

    expect(oraclePrices.length).toEqual(2);
    expect(oraclePrices[0]).toEqual(expect.objectContaining(defaultOraclePrice));
    expect(oraclePrices[1]).toEqual(expect.objectContaining({
      ...defaultOraclePrice,
      effectiveAtHeight: updatedHeight,
    }));
  });

  it('Successfully finds OraclePrice with effectiveAtHeight', async () => {
    await OraclePriceTable.create(defaultOraclePrice);

    const oraclePrices: OraclePriceFromDatabase[] = await OraclePriceTable.findAll(
      {
        effectiveAtHeight: defaultOraclePrice.effectiveAtHeight,
      },
      [],
      { readReplica: true },
    );

    expect(oraclePrices.length).toEqual(1);
    expect(oraclePrices[0]).toEqual(expect.objectContaining({
      ...defaultOraclePrice,
    }));
  });

  it('Successfully finds all OraclePrices effective before or after the height', async () => {
    await Promise.all([
      OraclePriceTable.create(defaultOraclePrice),
      OraclePriceTable.create({
        ...defaultOraclePrice,
        effectiveAtHeight: updatedHeight,
      }),
    ]);

    const oraclePrices: OraclePriceFromDatabase[] = await OraclePriceTable.findAll(
      {
        marketId: [defaultMarket.id],
        effectiveBeforeOrAtHeight: defaultOraclePrice.effectiveAtHeight,
      },
      [],
      {},
    );

    expect(oraclePrices.length).toEqual(1);
    expect(oraclePrices[0]).toEqual(expect.objectContaining(defaultOraclePrice));
  });

  it('Successfully finds all OraclePrices effective before or after time', async () => {
    const oraclePrice2: OraclePriceCreateObject = {
      ...defaultOraclePrice,
      effectiveAtHeight: updatedHeight,
      effectiveAt: '1982-05-25T00:00:00.000Z',
    };
    await Promise.all([
      OraclePriceTable.create(defaultOraclePrice),
      OraclePriceTable.create(oraclePrice2),
    ]);

    const oraclePrices: OraclePriceFromDatabase[] = await OraclePriceTable.findAll(
      {
        marketId: [defaultMarket.id],
        effectiveBeforeOrAt: '2000-05-25T00:00:00.000Z',
      },
      [],
      {},
    );

    expect(oraclePrices.length).toEqual(1);
    expect(oraclePrices[0]).toEqual(expect.objectContaining(oraclePrice2));
  });

  it('Successfully finds an OraclePrice', async () => {
    await OraclePriceTable.create(defaultOraclePrice);

    const oraclePrice: OraclePriceFromDatabase | undefined = await
    OraclePriceTable.findById(defaultOraclePriceId) as OraclePriceFromDatabase;
    expect(oraclePrice).toEqual(expect.objectContaining(defaultOraclePrice));
  });

  it('Successfully finds oracle prices in reverse chronological order by market id', async () => {
    const oraclePrice2: OraclePriceCreateObject = {
      ...defaultOraclePrice,
      effectiveAtHeight: updatedHeight,
    };
    await Promise.all([
      OraclePriceTable.create(defaultOraclePrice),
      OraclePriceTable.create(oraclePrice2),
    ]);

    const oraclePrices: OraclePriceFromDatabase[] = await OraclePriceTable
      .findOraclePricesInReverseChronologicalOrder(
        defaultMarket.id,
      ) as OraclePriceFromDatabase[];

    expect(oraclePrices.length).toEqual(2);
    expect(oraclePrices[0]).toEqual(expect.objectContaining(oraclePrice2));
    expect(oraclePrices[1]).toEqual(expect.objectContaining(defaultOraclePrice));
  });

  it('Successfully finds latest oracle price for market id', async () => {
    const oraclePrice2: OraclePriceCreateObject = {
      ...defaultOraclePrice,
      effectiveAtHeight: updatedHeight,
      effectiveAt: '1982-05-25T00:00:00.000Z',
    };
    await Promise.all([
      OraclePriceTable.create(defaultOraclePrice),
      OraclePriceTable.create(oraclePrice2),
    ]);

    const oraclePrice: OraclePriceFromDatabase = await OraclePriceTable
      .findMostRecentMarketOraclePrice(
        defaultMarket.id,
      ) as OraclePriceFromDatabase;

    expect(oraclePrice).toEqual(expect.objectContaining(oraclePrice2));
  });

  it('Successfully finds latest prices by effectiveAtHeight', async () => {
    const oraclePrice2: OraclePriceCreateObject = {
      ...defaultOraclePrice,
      spotPrice: '10000.05',
      pnlPrice: '10000.05',
      effectiveAtHeight: updatedHeight,
      effectiveAt: '1982-05-25T00:00:00.000Z',
    };
    await Promise.all([
      OraclePriceTable.create(defaultOraclePrice),
      OraclePriceTable.create(oraclePrice2),
      OraclePriceTable.create(defaultOraclePrice2),
    ]);

    const oraclePrices: PriceMap = await OraclePriceTable
      .findLatestPrices(
        updatedHeight,
      );

    expect(oraclePrices).toEqual(expect.objectContaining({
      [defaultOraclePrice.marketId]: {
        spotPrice: oraclePrice2.spotPrice,
        pnlPrice: oraclePrice2.pnlPrice,
      },
      [defaultOraclePrice2.marketId]: {
        spotPrice: defaultOraclePrice2.spotPrice,
        pnlPrice: defaultOraclePrice2.pnlPrice,
      },
    }));
  });

  it('Successfully finds the latest price and the price 24h ago', async () => {
    const now: string = DateTime.local().toISO();
    const lessThan24HAgo: string = DateTime.local().minus({ hour: 23 }).toISO();
    const moreThan24HAgo: string = DateTime.local().minus({ hour: 24, minute: 5 }).toISO();
    const wayMoreThan24HAgo: string = DateTime.local().minus({ hour: 25 }).toISO();

    const oraclePrice3: OraclePriceCreateObject = {
      ...defaultOraclePrice,
      spotPrice: '3',
      pnlPrice: '3',
      effectiveAtHeight: '3',
      effectiveAt: lessThan24HAgo,
    };
    const oraclePrice4: OraclePriceCreateObject = {
      ...defaultOraclePrice,
      spotPrice: '4',
      pnlPrice: '4',
      effectiveAtHeight: '4',
      effectiveAt: moreThan24HAgo,
    };
    const oraclePrice5: OraclePriceCreateObject = {
      ...defaultOraclePrice,
      spotPrice: '5',
      pnlPrice: '5',
      effectiveAtHeight: '5',
      effectiveAt: wayMoreThan24HAgo,
    };
    const oraclePrice6: OraclePriceCreateObject = {
      ...defaultOraclePrice,
      spotPrice: '6',
      pnlPrice: '6',
      effectiveAtHeight: '6',
      effectiveAt: now,
    };
    const oraclePrice7: OraclePriceCreateObject = {
      ...defaultOraclePrice,
      marketId: defaultMarket2.id,
      spotPrice: '7',
      pnlPrice: '7',
      effectiveAtHeight: '7',
      effectiveAt: lessThan24HAgo,
    };

    const blockHeights = ['3', '4', '6', '7'];

    const blockPromises = blockHeights.map((height) => BlockTable.create({
      ...defaultBlock,
      blockHeight: height,
    }),
    );

    await Promise.all(blockPromises);
    await Promise.all([
      OraclePriceTable.create(oraclePrice3),
      OraclePriceTable.create(oraclePrice4),
      OraclePriceTable.create(oraclePrice5),
      OraclePriceTable.create(oraclePrice6),
      OraclePriceTable.create(oraclePrice7),
    ]);

    const oraclePricesFrom24hAgo: PriceMap = await OraclePriceTable
      .getPricesFrom24hAgo();

    expect(oraclePricesFrom24hAgo).toEqual(expect.objectContaining({
      [defaultOraclePrice.marketId]: {
        spotPrice: oraclePrice4.spotPrice,
        pnlPrice: oraclePrice4.pnlPrice,
      },
    }));

    const latestPrices: PriceMap = await OraclePriceTable
      .getLatestPrices();

    expect(latestPrices).toEqual(expect.objectContaining({
      [defaultOraclePrice.marketId]: {
        spotPrice: oraclePrice6.spotPrice,
        pnlPrice: oraclePrice6.pnlPrice,
      },
      [defaultMarket2.id]: {
        spotPrice: oraclePrice7.spotPrice,
        pnlPrice: oraclePrice7.pnlPrice,
      },
    }));
  });

  it('Successfully finds latest prices respecting effectiveAtHeight', async () => {
    const oraclePrice2: OraclePriceCreateObject = {
      ...defaultOraclePrice,
      spotPrice: '10000.05',
      pnlPrice: '10000.05',
      effectiveAtHeight: updatedHeight,
      effectiveAt: '1982-05-25T00:00:00.000Z',
    };
    await Promise.all([
      OraclePriceTable.create(defaultOraclePrice),
      OraclePriceTable.create(oraclePrice2),
    ]);

    const oraclePrices: PriceMap = await OraclePriceTable
      .findLatestPrices(
        defaultOraclePrice.effectiveAtHeight,
      );

    expect(oraclePrices).toEqual(expect.objectContaining({
      [defaultOraclePrice.marketId]: {
        spotPrice: defaultOraclePrice.spotPrice,
        pnlPrice: defaultOraclePrice.pnlPrice,
      },
    }));
  });
});
