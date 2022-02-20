import {ISearchForm} from "../dto/ISearchForm";
import {ISearchResults} from "../dto/ISearchResult";
import baseApiUrl from "../const/constants";

const httpSnmpRequest = (request: ISearchForm): Promise<ISearchResults> => {
    const url = baseApiUrl + "/snmp/" + request.method

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
