package dtc



type DTCCode uint64

const (
	DTCCode_Unknown DTCCode = 0
	// byte 1
	//DTCCode_RCD_ACDC DTCCode = 0xff0601
	//DTCCode_RCD_ACDC_2       DTCCode = 0xff0603
	DTCCode_RCD_SELF_TEST    DTCCode = 0xff0601
	DTCCode_HIGH_TEMPERATURE DTCCode = 0xff0705
	Dtcerror_LOW_TEMPERATURE DTCCode = 0xff0707
	//DTCCode_MORE_TEMPERATURE=Reserved
	DTCCode_TEMPERATURE_OVERLOAD DTCCode = 0xff0701
	DTCCode_PE                   DTCCode = 0xff1002
	DTCCode_CP                   DTCCode = 0xff0801

	// byte 2
	DTCCode_PWM                          DTCCode = 0xff0F01
	DTCCode_ONE_LEVEL_CURRENT_OVERLOAD   DTCCode = 0xff0D03
	DTCCode_TWO_LEVEL_CURRENT_OVERLOAD   DTCCode = 0xff0D07
	DTCCode_ONE_LEVEL_VOLTAGE_OVERLOAD   DTCCode = 0xff0D05
	DTCCode_VOLTAGE_OWE_REDUCE_FREQUENCY DTCCode = 0xff0D11
	DTCCode_VOLTAGE_OWE                  DTCCode = 0xff0D02
	DTCCode_RELAY_SELF_TEST              DTCCode = 0xff0202
	DTCCode_RELAY_STATUS                 DTCCode = 0xff0201

	// byte 3
	DTCCode_RELAY_CONTROL                    DTCCode = 0xff0203
	DTCCode_WATCH_DOG_MCU_DEADLOCK           DTCCode = 0xff1104
	DTCCode_SUPERVISOR_KEY_MOLDUE_DEADLOCK   DTCCode = 0xff1101
	DTCCode_SUPERVISOR_UNKEY_MOLDUE_DEADLOCK DTCCode = 0xff1102
	DTCCode_DIP_VALUE                        DTCCode = 0xff0901
	DTCCode_PERSISTENCE_MEMORY_INCONSISTENCY DTCCode = 0xff1501
	DTCCode_PMIC_WATCHDOG_MESSAGE            DTCCode = 0xff0102

	// byte 4
	DTCCode_ADC_INIT       DTCCode = 0xff0802
	DTCCode_RFID_MOUDLE    DTCCode = 0xff0B01
	DTCCode_BLE_MOUDLE     DTCCode = 0xff0301
	DTCCode_CP_12V_MONITOR DTCCode = 0xff0804

	// byte 5
	DTCCode_TEMPERATURE_SENSOR_OPEN_CIRCUIT  DTCCode = 0xff0702
	DTCCode_TEMPERATURE_SENSOR_DRIFT_CIRCUIT DTCCode = 0xff0704
	DTCCode_SHORT_CIRCUIT                    DTCCode = 0xff0206
	DTCCode_EMERGENCY_STOP                   DTCCode = 0xff0204
	DTCCode_SHORT_CIRCUIT_CHARGING           DTCCode = 0xff0D09
	DTCCode_LONG_TIME_VOLTAGE_OVERLOAD       DTCCode = 0xff0D10
	DTCCode_LONG_TIME_VOLTAGE_OWE            DTCCode = 0xff0D12

	//DTCCode_TEMPERATURE_SENSOR_SHORT_CIRCUIT=
	//DTCCode_TEMPERATURE_SENSOR_OVERRANGE=Reserved
	//DTCCode_MCU_START_MBIST=Reserved

	// byte 6
	DTCCode_ELECTRIC_FREQUENCY                DTCCode = 0xff0D14
	DTCCode_CONNECT_CPO                       DTCCode = 0xff0E02
	DTCCode_CP_BROCKEN                        DTCCode = 0xff0803
	DTCCode_TMULTIPLE                         DTCCode = 0xff0706
	DTCCode_TWO_LEVEL_VOLTAGE_OVERLOAD        DTCCode = 0xff0D01
	DTCCode_UMULTIPLE                         DTCCode = 0xff0D06
	DTCCode_ONE_LEVEL_CURRENT_OVERLOAD_TWICE  DTCCode = 0xff0D13
	DTCCode_RELAY_PRE_DRIVER_HIGH_TEMPERATURE DTCCode = 0xff0F02

	// byte 7
	DTCCode_RELAY_PRE_DRIVER_OVERLOAD     DTCCode = 0xff0F03
	DTCCode_REGISTER                      DTCCode = 0xff0E01
	DTCCode_RCD_AC_LEAK                   DTCCode = 0xff0602
	DTCCode_RCD_DC_LEAK                   DTCCode = 0xff0603
	DTCCode_L2_ONE_LEVEL_VOLTAGE_OVERLOAD DTCCode = 0xff0DB1
	DTCCode_L2_TWO_LEVEL_VOLTAGE_OVERLOAD DTCCode = 0xff0DB2
	DTCCode_L2_LONG_TIME_VOLTAGE_OVERLOAD DTCCode = 0xff0DB3
	DTCCode_L2_VOLTAGE_OWE                DTCCode = 0xff0DB4

	// byte 8
	DTCCode_L2_VOLTAGE_OWE_REDUCE_FREQUENCY     DTCCode = 0xff0DB5
	DTCCode_L2_LONG_TIME_VOLTAGE_OWE            DTCCode = 0xff0DB6
	DTCCode_L2_ONE_LEVEL_CURRENT_OVERLOAD       DTCCode = 0xff0DB7
	DTCCode_L2_ONE_LEVEL_CURRENT_OVERLOAD_TWICE DTCCode = 0xff0DB8
	DTCCode_L2_TWO_LEVEL_CURRENT_OVERLOAD       DTCCode = 0xff0DB9
	DTCCode_L3_ONE_LEVEL_VOLTAGE_OVERLOAD       DTCCode = 0xff0DC1
	DTCCode_L3_TWO_LEVEL_VOLTAGE_OVERLOAD       DTCCode = 0xff0DC2
	DTCCode_L3_LONG_TIME_VOLTAGE_OVERLOAD       DTCCode = 0xff0DC3

	// byte 9
	DTCCode_L3_VOLTAGE_OWE                      DTCCode = 0xff0DC4
	DTCCode_L3_VOLTAGE_OWE_REDUCE_FREQUENCY     DTCCode = 0xff0DC5
	DTCCode_L3_LONG_TIME_VOLTAGE_OWE            DTCCode = 0xff0DC6
	DTCCode_L3_ONE_LEVEL_CURRENT_OVERLOAD       DTCCode = 0xff0DC7
	DTCCode_L3_ONE_LEVEL_CURRENT_OVERLOAD_TWICE DTCCode = 0xff0DC8
	DTCCode_L3_TWO_LEVEL_CURRENT_OVERLOAD       DTCCode = 0xff0DC9
	DTCCode_COM_BOARD_OFFLINE                   DTCCode = 0xff0e07
	DTCCode_GPS_ANTENNA                         DTCCode = 0xff0e03

	// byte 10
	DTCCode_4G_ANTENNA      DTCCode = 0xff0e04
	DTCCode_SIM_CARD        DTCCode = 0xff0e05
	DTCCode_COM_GENERAL     DTCCode = 0xff0e06
	DTCCode_SN_UNSET        DTCCode = 0xff1a01
	DTCCode_PARAMETER_UNSET DTCCode = 0xff1A02
)

type DTC struct {
	Desc  string
	Level DTCLevel
}
type DTCLevel int

const (
	DTCLevel_Danger   DTCLevel = 0
	DTCLevel_Hardware DTCLevel = 1
	DTCLevel_System   DTCLevel = 2
	DTCLevel_Warning  DTCLevel = 3
	DTCLevel_Notice   DTCLevel = 4
	DTCLevel_Info     DTCLevel = 5
)

func ToDtcCode(inCode Warnings) DTCCode {
	switch inCode {

	//byte 1
	case Warnings_CCM_RCDSelfTest:
		return DTCCode_RCD_SELF_TEST
	case Warnings_CCM_HighTemperature:
		return DTCCode_HIGH_TEMPERATURE
	case Warnings_CCM_LowTemperature:
		return Dtcerror_LOW_TEMPERATURE
	case Warnings_CCM_OverTemp:
		return DTCCode_TEMPERATURE_OVERLOAD
	case Warnings_CCM_PEFault:
		return DTCCode_PE
	case Warnings_CCM_CPFault:
		return DTCCode_CP

	// byte 2
	case Warnings_CCM_Pwm:
		return DTCCode_PWM
	case Warnings_CCM_OneLevelCurrentOverload:
		return DTCCode_ONE_LEVEL_CURRENT_OVERLOAD
	case Warnings_CCM_TwoLevelCurrentOverload:
		return DTCCode_TWO_LEVEL_CURRENT_OVERLOAD
	case Warnings_PSM_OneLevelVoltageOverload:
		return DTCCode_ONE_LEVEL_VOLTAGE_OVERLOAD
	case Warnings_PSM_VoltageOweReduceFrequency:
		return DTCCode_VOLTAGE_OWE_REDUCE_FREQUENCY
	case Warnings_PSM_VoltageOwe:
		return DTCCode_VOLTAGE_OWE
	case Warnings_CCM_RelaySelfTest:
		return DTCCode_RELAY_SELF_TEST
	case Warnings_CCM_RelayStatus:
		return DTCCode_RELAY_STATUS

	// byte 3
	case Warnings_CCM_RelayControl:
		return DTCCode_RELAY_CONTROL
	case Warnings_CCM_WatchDocMCUDeadlock:
		return DTCCode_WATCH_DOG_MCU_DEADLOCK
	case Warnings_CCM_SupervisorKeyModuleDeadlock:
		return DTCCode_SUPERVISOR_KEY_MOLDUE_DEADLOCK
	case Warnings_CCM_SupervisorUnKeyModuleDeadlock:
		return DTCCode_SUPERVISOR_UNKEY_MOLDUE_DEADLOCK
	case Warnings_CCM_DipValue:
		return DTCCode_DIP_VALUE
	case Warnings_SYS_PersistenceMemoryInconsistency:
		return DTCCode_PERSISTENCE_MEMORY_INCONSISTENCY
	case Warnings_CCM_PmicWatchDogMessage:
		return DTCCode_PMIC_WATCHDOG_MESSAGE

		// byte 4
	case Warnings_SYS_ADCInit:
		return DTCCode_ADC_INIT
	case Warnings_CCM_RFIDModule:
		return DTCCode_RFID_MOUDLE
	case Warnings_CCM_BLEModule:
		return DTCCode_BLE_MOUDLE
	case Warnings_CCM_CP12VMonitor:
		return DTCCode_CP_12V_MONITOR

	// byte 5
	case Warnings_CCM_TemperatureSensorOpenCircuit:
		return DTCCode_TEMPERATURE_SENSOR_OPEN_CIRCUIT
	case Warnings_CCM_TemperatureSensorDriftCircuit:
		return DTCCode_TEMPERATURE_SENSOR_DRIFT_CIRCUIT
	case Warnings_CCM_ShortCircuit:
		return DTCCode_SHORT_CIRCUIT
	case Warnings_CCM_EmergencyFault:
		return DTCCode_EMERGENCY_STOP
	case Warnings_CCM_ShortCircuitCharging:
		return DTCCode_SHORT_CIRCUIT_CHARGING
	case Warnings_PSM_LongTimeVoltageOverload:
		return DTCCode_LONG_TIME_VOLTAGE_OVERLOAD
	case Warnings_PSM_LongTimeVoltageOwe:
		return DTCCode_LONG_TIME_VOLTAGE_OWE

	// byte 6
	case Warnings_PSM_ElectricFrequency:
		return DTCCode_ELECTRIC_FREQUENCY
	case Warnings_SYS_ConnectCPO:
		return DTCCode_CONNECT_CPO
	case Warnings_CCM_CPBroken:
		return DTCCode_CP_BROCKEN
	case Warnings_CCM_MultipleTemperatureOverload:
		return DTCCode_TMULTIPLE
	case Warnings_PSM_TwoLevelVoltageOverload:
		return DTCCode_TWO_LEVEL_VOLTAGE_OVERLOAD
	case Warnings_PSM_MultipleVoltageException:
		return DTCCode_UMULTIPLE
	case Warnings_CCM_OneLevelCurrentOverloadTwice:
		return DTCCode_ONE_LEVEL_CURRENT_OVERLOAD_TWICE
	case Warnings_CCM_RelayPreDriverHighTemperature:
		return DTCCode_RELAY_PRE_DRIVER_HIGH_TEMPERATURE

	// byte 7
	case Warnings_CCM_RelayPreDriverOverload:
		return DTCCode_RELAY_PRE_DRIVER_OVERLOAD
	case Warnings_SYS_Register:
		return DTCCode_REGISTER
	case Warnings_PSM_RCD_AC_LEAK:
		return DTCCode_RCD_AC_LEAK
	case Warnings_PSM_RCD_DC_LEAK:
		return DTCCode_RCD_DC_LEAK
	case Warnings_CCM_L2OneLevelVoltageOverload:
		return DTCCode_L2_ONE_LEVEL_VOLTAGE_OVERLOAD
	case Warnings_CCM_L2TwoLevelVoltageOverload:
		return DTCCode_L2_TWO_LEVEL_VOLTAGE_OVERLOAD
	case Warnings_CCM_L2LongTimeVoltageOverload:
		return DTCCode_L2_LONG_TIME_VOLTAGE_OVERLOAD
	case Warnings_CCM_L2VoltageOwe:
		return DTCCode_L2_VOLTAGE_OWE

	// byte 8
	case Warnings_CCM_L2VoltageOweReduceFrequency:
		return DTCCode_L2_VOLTAGE_OWE_REDUCE_FREQUENCY
	case Warnings_CCM_L2LongTimeVoltageOwe:
		return DTCCode_L2_LONG_TIME_VOLTAGE_OWE
	case Warnings_CCM_L2OneLevelCurrentOverload:
		return DTCCode_L2_ONE_LEVEL_CURRENT_OVERLOAD
	case Warnings_CCM_L2OneLevelCurrentOverloadTwice:
		return DTCCode_L2_ONE_LEVEL_CURRENT_OVERLOAD_TWICE
	case Warnings_CCM_L2TwoLevelCurrentOverload:
		return DTCCode_L2_TWO_LEVEL_CURRENT_OVERLOAD
	case Warnings_CCM_L3OneLevelVoltageOverload:
		return DTCCode_L3_ONE_LEVEL_VOLTAGE_OVERLOAD
	case Warnings_CCM_L3TwoLevelVoltageOverload:
		return DTCCode_L3_TWO_LEVEL_VOLTAGE_OVERLOAD
	case Warnings_CCM_L3LongTimeVoltageOverload:
		return DTCCode_L3_LONG_TIME_VOLTAGE_OVERLOAD

	// byte 9
	case Warnings_CCM_L3VoltageOwe:
		return DTCCode_L3_VOLTAGE_OWE
	case Warnings_CCM_L3VoltageOweReduceFrequency:
		return DTCCode_L3_VOLTAGE_OWE_REDUCE_FREQUENCY
	case Warnings_CCM_L3LongTimeVoltageOwe:
		return DTCCode_L3_LONG_TIME_VOLTAGE_OWE
	case Warnings_CCM_L3OneLevelCurrentOverload:
		return DTCCode_L3_ONE_LEVEL_CURRENT_OVERLOAD
	case Warnings_CCM_L3OneLevelCurrentOverloadTwice:
		return DTCCode_L3_ONE_LEVEL_CURRENT_OVERLOAD_TWICE
	case Warnings_CCM_L3TwoLevelCurrentOverload:
		return DTCCode_L3_TWO_LEVEL_CURRENT_OVERLOAD
	case Warnings_SYS_ComBoardOffline:
		return DTCCode_COM_BOARD_OFFLINE
	case Warnings_SYS_GPSAntenna:
		return DTCCode_GPS_ANTENNA

	// byte 10
	case Warnings_SYS_4GAntenna:
		return DTCCode_4G_ANTENNA
	case Warnings_SYS_SimFault:
		return DTCCode_SIM_CARD
	case Warnings_SYS_ComGeneral:
		return DTCCode_COM_GENERAL
	case Warnings_SYS_SNUnset:
		return DTCCode_SN_UNSET
	case Warnings_SYS_ParameterUnset:
		return DTCCode_PARAMETER_UNSET
	}
	return DTCCode_Unknown
}

func ToCoreWarningCode(inCode DTCCode) Warnings {
	switch inCode {

	//byte 1
	case DTCCode_RCD_SELF_TEST:
		return Warnings_CCM_RCDSelfTest
	case DTCCode_HIGH_TEMPERATURE:
		return Warnings_CCM_HighTemperature
	case Dtcerror_LOW_TEMPERATURE:
		return Warnings_CCM_LowTemperature
	case DTCCode_TEMPERATURE_OVERLOAD:
		return Warnings_CCM_OverTemp
	case DTCCode_PE:
		return Warnings_CCM_PEFault
	case DTCCode_CP:
		return Warnings_CCM_CPFault

		// byte 2
	case DTCCode_PWM:
		return Warnings_CCM_Pwm
	case DTCCode_ONE_LEVEL_CURRENT_OVERLOAD:
		return Warnings_CCM_OneLevelCurrentOverload
	case DTCCode_TWO_LEVEL_CURRENT_OVERLOAD:
		return Warnings_CCM_TwoLevelCurrentOverload
	case DTCCode_ONE_LEVEL_VOLTAGE_OVERLOAD:
		return Warnings_PSM_OneLevelVoltageOverload
	case DTCCode_VOLTAGE_OWE_REDUCE_FREQUENCY:
		return Warnings_PSM_VoltageOweReduceFrequency
	case DTCCode_VOLTAGE_OWE:
		return Warnings_PSM_VoltageOwe
	case DTCCode_RELAY_SELF_TEST:
		return Warnings_CCM_RelaySelfTest
	case DTCCode_RELAY_STATUS:
		return Warnings_CCM_RelayStatus

		// byte 3
	case DTCCode_RELAY_CONTROL:
		return Warnings_CCM_RelayControl
	case DTCCode_WATCH_DOG_MCU_DEADLOCK:
		return Warnings_CCM_WatchDocMCUDeadlock
	case DTCCode_SUPERVISOR_KEY_MOLDUE_DEADLOCK:
		return Warnings_CCM_SupervisorKeyModuleDeadlock
	case DTCCode_SUPERVISOR_UNKEY_MOLDUE_DEADLOCK:
		return Warnings_CCM_SupervisorUnKeyModuleDeadlock
	case DTCCode_DIP_VALUE:
		return Warnings_CCM_DipValue
	case DTCCode_PERSISTENCE_MEMORY_INCONSISTENCY:
		return Warnings_SYS_PersistenceMemoryInconsistency
	case DTCCode_PMIC_WATCHDOG_MESSAGE:
		return Warnings_CCM_PmicWatchDogMessage

		// byte 4
	case DTCCode_ADC_INIT:
		return Warnings_SYS_ADCInit
	case DTCCode_RFID_MOUDLE:
		return Warnings_CCM_RFIDModule
	case DTCCode_BLE_MOUDLE:
		return Warnings_CCM_BLEModule
	case DTCCode_CP_12V_MONITOR:
		return Warnings_CCM_CP12VMonitor

		// byte 5
	case DTCCode_TEMPERATURE_SENSOR_OPEN_CIRCUIT:
		return Warnings_CCM_TemperatureSensorOpenCircuit
	case DTCCode_TEMPERATURE_SENSOR_DRIFT_CIRCUIT:
		return Warnings_CCM_TemperatureSensorDriftCircuit
	case DTCCode_SHORT_CIRCUIT:
		return Warnings_CCM_ShortCircuit
	case DTCCode_EMERGENCY_STOP:
		return Warnings_CCM_EmergencyFault
	case DTCCode_SHORT_CIRCUIT_CHARGING:
		return Warnings_CCM_ShortCircuitCharging
	case DTCCode_LONG_TIME_VOLTAGE_OVERLOAD:
		return Warnings_PSM_LongTimeVoltageOverload
	case DTCCode_LONG_TIME_VOLTAGE_OWE:
		return Warnings_PSM_LongTimeVoltageOwe

		// byte 6
	case DTCCode_ELECTRIC_FREQUENCY:
		return Warnings_PSM_ElectricFrequency
	case DTCCode_CONNECT_CPO:
		return Warnings_SYS_ConnectCPO
	case DTCCode_CP_BROCKEN:
		return Warnings_CCM_CPBroken
	case DTCCode_TMULTIPLE:
		return Warnings_CCM_MultipleTemperatureOverload
	case DTCCode_TWO_LEVEL_VOLTAGE_OVERLOAD:
		return Warnings_PSM_TwoLevelVoltageOverload
	case DTCCode_UMULTIPLE:
		return Warnings_PSM_MultipleVoltageException
	case DTCCode_ONE_LEVEL_CURRENT_OVERLOAD_TWICE:
		return Warnings_CCM_OneLevelCurrentOverloadTwice
	case DTCCode_RELAY_PRE_DRIVER_HIGH_TEMPERATURE:
		return Warnings_CCM_RelayPreDriverHighTemperature

		// byte 7
	case DTCCode_RELAY_PRE_DRIVER_OVERLOAD:
		return Warnings_CCM_RelayPreDriverOverload
	case DTCCode_REGISTER:
		return Warnings_SYS_Register
	case DTCCode_RCD_AC_LEAK:
		return Warnings_PSM_RCD_AC_LEAK
	case DTCCode_RCD_DC_LEAK:
		return Warnings_PSM_RCD_DC_LEAK
	case DTCCode_L2_ONE_LEVEL_VOLTAGE_OVERLOAD:
		return Warnings_CCM_L2OneLevelVoltageOverload
	case DTCCode_L2_TWO_LEVEL_VOLTAGE_OVERLOAD:
		return Warnings_CCM_L2TwoLevelVoltageOverload
	case DTCCode_L2_LONG_TIME_VOLTAGE_OVERLOAD:
		return Warnings_CCM_L2LongTimeVoltageOverload
	case DTCCode_L2_VOLTAGE_OWE:
		return Warnings_CCM_L2VoltageOwe

		// byte 8
	case DTCCode_L2_VOLTAGE_OWE_REDUCE_FREQUENCY:
		return Warnings_CCM_L2VoltageOweReduceFrequency
	case DTCCode_L2_LONG_TIME_VOLTAGE_OWE:
		return Warnings_CCM_L2LongTimeVoltageOwe
	case DTCCode_L2_ONE_LEVEL_CURRENT_OVERLOAD:
		return Warnings_CCM_L2OneLevelCurrentOverload
	case DTCCode_L2_ONE_LEVEL_CURRENT_OVERLOAD_TWICE:
		return Warnings_CCM_L2OneLevelCurrentOverloadTwice
	case DTCCode_L2_TWO_LEVEL_CURRENT_OVERLOAD:
		return Warnings_CCM_L2TwoLevelCurrentOverload
	case DTCCode_L3_ONE_LEVEL_VOLTAGE_OVERLOAD:
		return Warnings_CCM_L3OneLevelVoltageOverload
	case DTCCode_L3_TWO_LEVEL_VOLTAGE_OVERLOAD:
		return Warnings_CCM_L3TwoLevelVoltageOverload
	case DTCCode_L3_LONG_TIME_VOLTAGE_OVERLOAD:
		return Warnings_CCM_L3LongTimeVoltageOverload

		// byte 9
	case DTCCode_L3_VOLTAGE_OWE:
		return Warnings_CCM_L3VoltageOwe
	case DTCCode_L3_VOLTAGE_OWE_REDUCE_FREQUENCY:
		return Warnings_CCM_L3VoltageOweReduceFrequency
	case DTCCode_L3_LONG_TIME_VOLTAGE_OWE:
		return Warnings_CCM_L3LongTimeVoltageOwe
	case DTCCode_L3_ONE_LEVEL_CURRENT_OVERLOAD:
		return Warnings_CCM_L3OneLevelCurrentOverload
	case DTCCode_L3_ONE_LEVEL_CURRENT_OVERLOAD_TWICE:
		return Warnings_CCM_L3OneLevelCurrentOverloadTwice
	case DTCCode_L3_TWO_LEVEL_CURRENT_OVERLOAD:
		return Warnings_CCM_L3TwoLevelCurrentOverload
	case DTCCode_COM_BOARD_OFFLINE:
		return Warnings_SYS_ComBoardOffline
	case DTCCode_GPS_ANTENNA:
		return Warnings_SYS_GPSAntenna

		// byte 10
	case DTCCode_4G_ANTENNA:
		return Warnings_SYS_4GAntenna
	case DTCCode_SIM_CARD:
		return Warnings_SYS_SimFault
	case DTCCode_COM_GENERAL:
		return Warnings_SYS_ComGeneral
	case DTCCode_SN_UNSET:
		return Warnings_SYS_SNUnset
	case DTCCode_PARAMETER_UNSET:
		return Warnings_SYS_ParameterUnset
	}
	return Warnings_Other
}
