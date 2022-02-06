import {ISearchForm} from "../dto/ISearchForm";
import {ISearchResults} from "../dto/ISearchResult";

const httpSnmpRequest = (request: ISearchForm): Promise<ISearchResults> => {
    // const url = window.location.href + "snmp/" + request.method
    const url = "http://127.0.0.1:7000/snmp/" + request.method;

    const requestHeaders: HeadersInit = new Headers();
    requestHeaders.set('Content-Type', 'application/json');

    return fetch(url, {
        method: "POST",
        headers: requestHeaders,
        body: JSON.stringify(request),
        keepalive: true,
        mode: "cors"
    }).then(result => result.json());
}

export default httpSnmpRequest;
