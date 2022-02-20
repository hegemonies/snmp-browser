import {Text} from "@yandex/ui/Text/bundle";
import './Header.sass'

const Header = () => {
    return (
        <header>
            <Text
                typography={"headline-xl"}
                color={"secondary"}
            >
                SNMP Browser
            </Text>
        </header>
    )
}

export default Header;
