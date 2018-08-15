package main

import (
    "fmt"
    "encoding/json"
)
const (
jsonStuff = `{
  "apPropertyBlock": {
    "bpWiredMacaddr": "00:04:96:97:E9:EE",
    "ruModel": "X460G2-24t-10G4",
    "ruSerialNumber": "1405G-00125",
    "ruSwVersion": "22.6.0.20"
  },
  "deviceInfo": {
    "firmwareRevision": "22.6.0.20",
    "license": {
      "effectiveLicense": "Advanced Edge",
      "enabledLicense": "Advanced Edge",
      "featurePacks": [
        "DirectAttach"
      ],
      "portCapacityLicense": ""
    },
    "macAddr": "00:04:96:97:E9:EE",
    "mgmtIpAddr": "10.68.67.93",
    "mgmtPort": "0",
    "operatingSystem": "ExtremeXOS",
    "ports": [
      {
        "portList": [
          "mgmt-1"
        ],
        "portSpeed": 100,
        "portType": "mgmt"
      },
      {
        "portList": [
          "25",
          "26",
          "27",
          "28"
        ],
        "portSpeed": 0,
        "portType": "access"
      },
      {
        "portList": [
          "1",
          "2",
          "3",
          "4",
          "5",
          "6",
          "7",
          "8",
          "9",
          "10",
          "11",
          "12",
          "13",
          "14",
          "15",
          "16",
          "17",
          "18",
          "19",
          "20",
          "21",
          "22",
          "23",
          "24"
        ],
        "portSpeed": 1000,
        "portType": "access"
      },
      {
        "portList": [
          "29",
          "30",
          "31",
          "32",
          "33",
          "34"
        ],
        "portSpeed": 10000,
        "portType": "uplink"
      }
    ],
    "serialNumber": "1405G-00125",
    "sysContact": "support@extremenetworks.com, +1 888 257 3000",
    "sysDescr": "ExtremeXOS (X460G2-24t-10G4) version 22.6.0.20 xos_22.6 by dhammers on Wed Aug 8 17:10:39 EDT 2018",
    "sysName": "X460G2-24t-10G4",
    "sysObjectID": "1.3.6.1.4.1.1916.2.197",
    "sysUpTime": 18829200
  }
}`
)

func main() {
    var jsonData map[string]interface{}
    _ = json.Unmarshal([]byte(jsonStuff) , &jsonData)

    // access the nested portList
    abc := jsonData["deviceInfo"].(map[string]interface{})
    def := abc["ports"].([]interface{})
    ghi := def[3].(map[string]interface{})
    jkl := ghi["portList"].([]interface{})

    // print some things we found
    fmt.Println("sysName", abc["sysName"])
    fmt.Println("serialNumber", abc["serialNumber"])
    fmt.Println("firmwareRevision", abc["firmwareRevision"])

    // print the port list we found
    for _, port := range jkl {
        fmt.Println(port)
    }

    // add some stuff to the end of the port list
    jkl = append(jkl, "a", "b", "c")

    // put it back into the original JSON structure
    ghi["portList"] = jkl

    // pretty print the JSON structure to see the changes
    bufx, _ := json.MarshalIndent(jsonData, "", "  ")
    fmt.Println(string(bufx))

    return 
}
