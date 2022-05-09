import {ISearchForm} from "../dto/ISearchForm";
import {ISearchResults} from "../dto/ISearchResult";
import baseApiUrl from "../const/constants";

const httpSnmpRequest = (request: ISearchForm): Promise<ISearchResults> => {

    let url = "";

    if (baseApiUrl[baseApiUrl.length - 1] === "/") {
        url = baseApiUrl + "snmp/" + request.method
    } else {
        url = baseApiUrl + "/snmp/" + request.method
    }

    console.log(url)

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
