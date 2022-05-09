export enum SnmpMethod {
    GET = "get",
    WALK = "walk"
}

export enum SnmpVersion {
    V1 = "1",
    V2C = "2c",
    V3 = "3"
}

export const snmpMethodDefault = SnmpMethod.GET;
export const targetHostnameDefault = "172.24.0.5";
export const oidsDefault = ["1.3.6.1.2.1.1.5.0"];
export const communitesDefault = ["public"];
export const snmpPortDefault = 161;
export const snmpVersionDefault = SnmpVersion.V2C;
export const snmpTimeoutDefault = 5; // sec
export const snmpRetriesDefault = 0;

let baseApiUrl = window.location.href;

if (process.env.REACT_APP_ENV === "dev") {
    baseApiUrl = 'http://' + process.env.REACT_APP_API_URL;
}

export default baseApiUrl;
