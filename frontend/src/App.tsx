import {useState} from 'react';
import './App.sass'

import Header from "./components/header/Header";
import SearchForm from "./components/form/SearchForm";
import SearchResult from "./components/result/SearchResult";
import {ISearchResult, ISearchResults} from "./dto/ISearchResult";
import { Divider } from '@yandex/ui/Divider';

const App = () => {
    const [searchResult, setSearchResult] = useState<ISearchResult[]>([]);
    const updateSearchResult = (results: ISearchResults) => {
      setSearchResult(results.results);
    }

    return (
        <div>
            <div className="app">
                <Header/>
                <Divider color={"#6768ab"} size={2}/>
                <div className="content">
                    <SearchForm onSubmit={updateSearchResult}/>
                    <SearchResult searchResults={searchResult}/>
                </div>
            </div>
        </div>
    )
}

export default App;
