import {
  CountryHeaders,
} from '@klyraprotocol-indexer/compliance';

import { IncomingMessage } from '../types';

export class CountryRestrictor {
  public getCountry(req: IncomingMessage): string | undefined {
    const countryHeaders: CountryHeaders = req.headers as CountryHeaders;
    return countryHeaders['cf-ipcountry'];
  }
}
