import {SnmpMethod, SnmpVersion} from "../const/constants";

export interface ISearchForm {
    method: SnmpMethod;
    targetHostname: string;
    oids: string[];
    communities: string[];
    port: number;
    version: SnmpVersion;
    timeout: number;
    retries: number;
}
