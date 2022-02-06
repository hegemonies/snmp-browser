import {ISearchResult} from "../../dto/ISearchResult";
import React from "react";
import Table from './Table';
import {Spacer} from "@yandex/ui/Spacer";

const SearchResult = (props: { searchResults: ISearchResult[] }) => {
    const columns = [{
        Header: "Name",
        accessor: "name"
    }, {
        Header: "Type",
        accessor: "type"
    }, {
        Header: "Value",
        accessor: "value"
    }];

    if (props.searchResults.length === 0) {
        return (<div/>);
    }

    return (
        <Spacer left={10}>
            <Table
                columns={columns}
                data = {props.searchResults}
            />
        </Spacer>
    );
}

export default SearchResult;
