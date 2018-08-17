
# interface definiation in python
cdef extern from "../c_src/change_patient_drv.c":
    void change_patient(int)

cdef extern from "../c_src/get_battery_level.c":
    float get_battery_level()

cdef extern from "../c_src/get_wifi_level.c":
    float get_wifi_level()


# interface implementation in python
cpdef changePatient(id):
    print "at python level: Change paitent:",id
    # call c function
    change_patient(id)

cpdef getBatteryLevel():
    # call c function
    batteryLevel = get_battery_level()
    return batteryLevel

cpdef getWifiLevel():
    # call c function
    batteryLevel = get_wifi_level()
    return batteryLevel