import React from 'react';
import {useTable} from "react-table"
import {ISearchResult} from "../../dto/ISearchResult";
import {Text} from "@yandex/ui/Text/bundle";

const Table = (props: {columns: any, data: ISearchResult[]}) => {
    const {
        getTableProps,
        getTableBodyProps,
        headerGroups,
        rows,
        prepareRow,
    } = useTable({columns: props.columns, data: props.data})

    return (
        <table {...getTableProps()}>
            <thead>
            {
                headerGroups.map(headerGroup => (
                    <tr {...headerGroup.getHeaderGroupProps()}>
                        {
                            headerGroup.headers.map( column => (
                                <th {...column.getHeaderProps()}>
                                    {
                                        <Text typography={"headline-m"} color={"secondary"}>
                                            {column.render('Header')}
                                        </Text>
                                    }
                                </th>
                            ))
                        }
                    </tr>
                ))
            }
            </thead>
            <tbody {...getTableBodyProps()}>
            { // loop over the rows
                rows.map(row => {
                    prepareRow(row)
                    return (
                        <tr {...row.getRowProps()}>
                            { // loop over the rows cells
                                row.cells.map(cell => (
                                    <td {...cell.getCellProps()}>
                                        <Text typography={"body-short-l"}>
                                            {cell.render('Cell')}
                                        </Text>
                                    </td>
                                ))
                            }
                        </tr>
                    )
                })
            }
            <tr>
                <td></td>
            </tr>
            </tbody>
        </table>
    );
}

export default Table;