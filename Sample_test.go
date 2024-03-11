package main

import (
    "fmt"
    "time"
)

// OspfTestcase struct represents the test case for OSPF configuration and validation
type OspfTestcase struct{}

// ConfiguringOspfTimers configures OSPF timers on routers
func (ot *OspfTestcase) ConfiguringOspfTimers(testbed Testbed) {
    fmt.Println("Configure OSPF timers on routers")
    Ospfexec(uut1, device1["intf1"])
    Ospfexec(uut3, device3["intf1"])
    time.Sleep(15 * time.Second)
}

// CheckPingOspfTimers checks if OSPF timers are configured and validates OSPF neighbors and routes
func (ot *OspfTestcase) CheckPingOspfTimers(testbed Testbed) {
    fmt.Println("Check OSPF timers configured or not")
    time.Sleep(30 * time.Second)
    ospfConfig := uut1.execute("show ip ospf interface ethernet 1/11")
    fmt.Println(ospfConfig)
    if strings.Contains(ospfConfig, "Timer intervals: Hello 30, Dead 100") {
        fmt.Println("OSPF timers configured on device successful")
    } else {
        fmt.Println("OSPF timers are not configured on device unsuccessful")
    }

    // Similar validations for OSPF neighbors and routes

    fmt.Println("Ping the IP configured on device2:", uut3.name, "to check inbound")
    for i := 0; i < 3; i++ {
        result1 := uut2.execute("ping " + device3["ip_address"])
        resDict := ValidatePing(result1)
        fmt.Println("++++++++++++++++++++++++++++")
        fmt.Println(resDict)
        fmt.Println("++++++++++++++++++++++++++++")

        if resDict["sent_pkt"] == resDict["receive_pkt"] && resDict["pkt_loss"] == "0.00%" {
            fmt.Printf("Sent: %s packets and received: %s packets and packet loss: %s\n", resDict["sent_pkt"], resDict["receive_pkt"], resDict["pkt_loss"])
            fmt.Println("Success: After applying for OSPF, ping got successful")
        } else {
            fmt.Printf("Sent: %s packets and received: %s packets and packet loss: %s\n", resDict["sent_pkt"], resDict["receive_pkt"], resDict["pkt_loss"])
            fmt.Println("Failed: After applying for OSPF, ping got failed")
        }
    }
}

// UnconfigureOspfTimersOnDevice unconfigures OSPF timers on routers
func (ot *OspfTestcase) UnconfigureOspfTimersOnDevice(testbed Testbed) {
    fmt.Println("Unconfigure OSPF timers on routers")
    CleaningOspf(uut1, device1["intf1"])
    CleaningOspf(uut3, device3["intf1"])
}

func main() {
    ospfTest := &OspfTestcase{}
    ospfTest.ConfiguringOspfTimers(testbed)
    ospfTest.CheckPingOspfTimers(testbed)
    ospfTest.UnconfigureOspfTimersOnDevice(testbed)
}
