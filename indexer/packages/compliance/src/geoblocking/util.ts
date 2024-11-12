import config from '../config';

const RESTRICTED_COUNTRY_CODES: Set<string> = new Set(config.RESTRICTED_COUNTRIES.split(','));
