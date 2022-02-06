const snmp = require("net-snmp");

const options = {
    port: 161,
    retries: 0,
    timeout: 5000,
    backoff: 0.0,
    transport: "udp4",
    version: snmp.Version2c,
    backwardsGetNexts: false,
    idBitsSize: 32
};

const session = snmp.createSession("10.24.16.199", "public", options);

function printVarbind(varbind) {
    if (snmp.isVarbindError(varbind)) {
        console.error(snmp.varbindError(varbind))
    } else {
        console.log(varbind.oid + " : " + varbind.type + " = " + varbind.value)
    }
}

function testGet() {
    const oids = ["1.3.6.1.2.1.1.5.0", "1.3.6.1.2.1.1.6.0"]

    session.get(oids, function (error, varbinds) {
        if (error) {
            console.error(error);
        } else {
            for (let i = 0; i < varbinds.length; i++) {
                printVarbind(varbinds[i])
            }
        }

        session.close()
    })
}

function testWalk() {
    const oid = "1.3.6.1.2.1.2.2.1.3"

    function doneCb(error) {
        if (error) {
            console.error(error);
        }
    }

    function feedCb(varbinds) {
        for (let i = 0; i < varbinds.length; i++) {
            printVarbind(varbinds[i])
        }
    }

    const maxRepetitions = 20;

    session.walk(oid, maxRepetitions, feedCb, doneCb);
}

testWalk()
