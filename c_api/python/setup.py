
from distutils.core import setup
from distutils.extension import Extension
from Cython.Distutils import build_ext
from Cython.Build.Dependencies import cythonize

ext_modules = [
        Extension("app", ["./src/app.py"]),
		Extension("DataPipeline", ["./src/DataPipeline.py"]),
		Extension("DeviceDrvSim", ["./src/DeviceDrvSim.py"]),
		Extension("HealthMonitor", ["./src/HealthMonitor.py"]),
		Extension("Processing", ["./src/Processing.py"]),
		Extension("RestServer", ["./src/RestServer.py"]),
		Extension("System", ["./src/System.py"]),
		Extension("bdmac", ["./src/bdmac.py"]),
		
		Extension("deviceDriver.ArmDevice", ["./src/deviceDriver/ArmDevice.py"]),
		Extension("deviceDriver.battery_level", ["./src/deviceDriver/battery_level.py"]),
		Extension("deviceDriver.cell_signal", ["./src/deviceDriver/cell_signal.py"]),
		Extension("deviceDriver.change_patient", ["./src/deviceDriver/change_patient.py"]),
		Extension("deviceDriver.disk_usage_level", ["./src/deviceDriver/disk_usage_level.py"]),
		Extension("deviceDriver.electronodes", ["./src/deviceDriver/electronodes.py"]),
		Extension("deviceDriver.flash_disk_connected", ["./src/deviceDriver/flash_disk_connected.py"]),
		Extension("deviceDriver.local_socket", ["./src/deviceDriver/local_socket.py"]),
		Extension("deviceDriver.usb_connected", ["./src/deviceDriver/usb_connected.py"]),
		Extension("deviceDriver.volume_control", ["./src/deviceDriver/volume_control.py"]),
		Extension("deviceDriver.wifi_level", ["./src/deviceDriver/wifi_level.py"]),

		Extension("algorithm_module.algo_api", ["./src/algorithm_module/algo_api.py"]),	
		Extension("algorithm_module.data", ["./src/algorithm_module/data.py"]),	
		Extension("algorithm_module.data_class", ["./src/algorithm_module/data_class.py"]),	
		Extension("algorithm_module.featurePointClass", ["./src/algorithm_module/featurePointClass.py"]),		
		Extension("algorithm_module._filter_class_", ["./src/algorithm_module/_filter_class_.py"]),	
		Extension("algorithm_module.systemParam_class", ["./src/algorithm_module/systemParam_class.py"]),		
		Extension("algorithm_module.VarClass", ["./src/algorithm_module/VarClass.py"]),
]

setup(
  name = 'Backend',
  cmdclass = {'build_ext': build_ext},
  ext_modules=cythonize(ext_modules, compiler_directives={'always_allow_keywords': True})
)
