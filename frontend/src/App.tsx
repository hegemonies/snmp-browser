import React, {useState} from 'react';
import './App.sass'

import Header from "./Header";
import SearchForm from "./components/form/SearchForm";
import SearchResult from "./components/result/SearchResult";
import {ISearchResult, ISearchResults} from "./dto/ISearchResult";

const App = () => {
    const [searchResult, setSearchResult] = useState<ISearchResult[]>([]);
    const updateSearchResult = (results: ISearchResults) => {
      setSearchResult(results.results);
    }

    return (
        <div className="app">
            <div className="app-row">
                <Header/>
                <SearchForm onSubmit={updateSearchResult}/>
            </div>
            <div className="app-row">
                <SearchResult searchResults={searchResult}/>
            </div>
        </div>
    )
}

export default App;
