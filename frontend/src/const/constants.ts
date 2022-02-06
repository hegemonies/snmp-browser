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
export const targetHostnameDefault = "10.24.16.69";
export const oidsDefault = ["1.3.6.1.2.1.1.5.0"];
export const communitesDefault = ["public"];
export const snmpPortDefault = 161;
export const snmpVersionDefault = SnmpVersion.V2C;
export const snmpTimeoutDefault = 5 // sec
export const snmpRetriesDefault = 0