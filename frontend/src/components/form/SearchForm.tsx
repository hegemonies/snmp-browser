import React, {useState} from "react";
import {Textinput} from "@yandex/ui/Textinput/desktop/bundle";
import {Button} from "@yandex/ui/Button/desktop/bundle";
import {ISearchForm} from "../../dto/ISearchForm";
import {
    communitesDefault,
    oidsDefault, SnmpMethod,
    snmpMethodDefault,
    snmpPortDefault, snmpRetriesDefault, snmpTimeoutDefault, SnmpVersion, snmpVersionDefault,
    targetHostnameDefault
} from "../../const/constants";
import './SearchForm.sass';
import {RadioButton} from '@yandex/ui/RadioButton/desktop/bundle'
import httpSnmpRequest from "../../api/api";
import {ISearchResults} from "../../dto/ISearchResult";
import {ToastContainer, toast} from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';
import {Text} from "@yandex/ui/Text/bundle";

const SearchForm = (props: { onSubmit: (results: ISearchResults) => void }) => {

    const [searchInProgress, setSearchInProgress] = useState(false);

    const errorNotification = (message: string) => {
        toast(message, {
            position: "top-right",
            autoClose: 3000,
            hideProgressBar: false,
            closeOnClick: true,
            pauseOnHover: true,
            draggable: true,
            progress: undefined,
            type: "error",
            theme: "light"
        });
    }

    const pressButton = () => {
        const button = document.getElementById("send-button");
        button?.click();
    }

    const sendRequest = () => {

        setSearchInProgress(true);

        if (value.targetHostname.length === 0 || !isIpValid(value.targetHostname)) {
            errorNotification("Target hostname is invalid");
            return;
        }

        if (value.oids.length === 0) {
            errorNotification("oids is invalid");
            return;
        }

        httpSnmpRequest(value).then(
            value => {
                props.onSubmit(value);
                setSearchInProgress(false);
            }, error => {
                errorNotification(error);
                setSearchInProgress(false);
            }
        );
    }

    const onKeyDown = (event: React.KeyboardEvent<HTMLDivElement>): void => {
        if (event.key === 'Enter') {
            event.preventDefault();
            event.stopPropagation();
            pressButton();
        }
    }

    const validateTargetHostname = (targetHostname: string): "error" | undefined => {

        if (targetHostname.length === 0) {
            return undefined
        }

        if (!isIpValid(targetHostname)) {
            return "error"
        }

        return undefined
    }

    const isIpValid = (ip: string): boolean => {
        return /^(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$/.test(ip);
    }

    const emptyForm: ISearchForm = {
        method: snmpMethodDefault,
        targetHostname: targetHostnameDefault,
        oids: oidsDefault,
        communities: communitesDefault,
        port: snmpPortDefault,
        version: snmpVersionDefault,
        timeout: snmpTimeoutDefault,
        retries: snmpRetriesDefault
    }

    const [value, setValue] = useState(emptyForm);

    return (
        <div className="search-form">

            <ToastContainer/>

            <div className="search-form-row">
                <Text typography={"control-m"} color={"secondary"}>Target hostname</Text>

                <Textinput
                    size="m"
                    view="default"
                    placeholder="192.168.0.1"
                    className="search-form-input"
                    value={value.targetHostname}
                    onChange={(event) => setValue({
                        ...value,
                        targetHostname: event.target.value
                    })}
                    onKeyDown={(event) => onKeyDown(event)}
                    required={true}
                    hasClear={true}
                    pressed={true}
                    state={validateTargetHostname(value.targetHostname)}
                />
            </div>
            <div className="search-form-row">
                <Text typography={"control-m"} color={"secondary"}>SNMP oids</Text>

                <Textinput
                    size="m"
                    view="default"
                    placeholder="1.3.6.1.2.1.2.2.1.10, 1.3.6.1.2.1.2.2.1.15"
                    className="search-form-input"
                    value={value.oids.join("")}
                    onChange={(event) => setValue({
                        ...value,
                        oids: event.target.value.split(",")
                    })}
                    onKeyDown={(event) => onKeyDown(event)}
                    required={true}
                    hasClear={true}
                />
            </div>
            <div className="search-form-row">
                <Text typography={"control-m"} color={"secondary"}>SNMP communities</Text>

                <Textinput
                    size="m"
                    view="default"
                    placeholder="public, private"
                    className="search-form-input"
                    value={value.communities.join("")}
                    onChange={(event) => setValue({
                        ...value,
                        communities: event.target.value.split(",")
                    })}
                    onKeyDown={(event) => onKeyDown(event)}
                    hasClear={true}
                />
            </div>
            <div className="search-form-row">
                <Text typography={"control-m"} color={"secondary"}>SNMP target port</Text>

                <Textinput
                    size="m"
                    view="default"
                    placeholder="161"
                    className="search-form-input"
                    value={value.port}
                    onChange={(event) => setValue({
                        ...value,
                        port: Number(event.target.value)
                    })}
                    onKeyDown={(event) => onKeyDown(event)}
                />
            </div>
            <div className="search-form-row">
                <RadioButton
                    size="m"
                    view="default"
                    className="search-form-input search-form-radio-button"
                    value={value.method}
                    options={[
                        {value: SnmpMethod.GET, children: "GET"},
                        {value: SnmpMethod.WALK, children: "WALK"},
                    ]}
                    onChange={(event) => setValue({
                        ...value,
                        method: event.target.value as SnmpMethod
                    })}
                />
                <RadioButton
                    size="m"
                    view="default"
                    className="search-form-input search-form-radio-button"
                    value={value.version}
                    options={[
                        {value: SnmpVersion.V1, children: "Version 1"},
                        {value: SnmpVersion.V2C, children: "Version 2c"},
                        {value: SnmpVersion.V3, children: "Version 3", disabled: true},
                    ]}
                    onChange={(event) => setValue({
                        ...value,
                        version: event.target.value as SnmpVersion
                    })}
                />
            </div>
            <div className="search-form-row">
                <Text typography={"control-m"} color={"secondary"}>SNMP request timeout in sec</Text>
                <Textinput
                    size="m"
                    view="default"
                    placeholder="5"
                    className="search-form-input"
                    value={value.timeout}
                    onChange={(event) => setValue({
                        ...value,
                        timeout: Number(event.target.value)
                    })}
                    onKeyDown={(event) => onKeyDown(event)}
                    inputMode={"numeric"}
                    type={"number"}
                    min={1}
                    max={120}
                    hasClear={true}
                />
            </div>
            <div className="search-form-row">
                <Text typography={"control-m"} color={"secondary"}>SNMP request retries</Text>
                <Textinput
                    size="m"
                    view="default"
                    placeholder="0"
                    className="search-form-input"
                    value={value.retries}
                    onChange={(event) => setValue({
                        ...value,
                        retries: Number(event.target.value)
                    })}
                    onKeyDown={(event) => onKeyDown(event)}
                    inputMode={"numeric"}
                    type={"number"}
                    min={0}
                    max={15}
                />
            </div>

            <div className="search-form-row">
                <Button
                    id="send-button"
                    view="default"
                    size="m"
                    theme={"action"}
                    width="auto"
                    className="search-form-button"
                    onClick={() => sendRequest()}
                    progress={searchInProgress}
                >
                    Search
                </Button>
            </div>
        </div>
    )
}

export default SearchForm;
