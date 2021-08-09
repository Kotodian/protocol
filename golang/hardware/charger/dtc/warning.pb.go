package dtc

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Warnings int32

const (
	//*系统故障
	Warnings_Other                              Warnings = 0
	Warnings_SYS_StorageFull                    Warnings = 100
	Warnings_SYS_4GFault                        Warnings = 101
	Warnings_SYS_SimFault                       Warnings = 102
	Warnings_SYS_DialFault                      Warnings = 103
	Warnings_SYS_ConnectError                   Warnings = 104
	Warnings_SYS_Disconnect                     Warnings = 105
	Warnings_SYS_PowerFailure                   Warnings = 106
	Warnings_SYS_MeterError                     Warnings = 107
	Warnings_SYS_CardReaderError                Warnings = 108
	Warnings_SYS_LCDError                       Warnings = 109
	Warnings_SYS_MeterMeasureException          Warnings = 110
	Warnings_SYS_MeterCollectsException         Warnings = 111
	Warnings_SYS_SoftFault                      Warnings = 112
	Warnings_SYS_AirInletOverTemp               Warnings = 113
	Warnings_SYS_AirInletUnderTemp              Warnings = 114
	Warnings_SYS_AirOutletOverTemp              Warnings = 115
	Warnings_SYS_SlaveFault                     Warnings = 116
	Warnings_SYS_StorageFailure                 Warnings = 117
	Warnings_SYS_PersistenceMemoryInconsistency Warnings = 118
	Warnings_SYS_ADCInit                        Warnings = 119
	Warnings_SYS_ConnectCPO                     Warnings = 120
	Warnings_SYS_Register                       Warnings = 121
	Warnings_SYS_ComBoardOffline                Warnings = 122
	Warnings_SYS_GPSAntenna                     Warnings = 123
	Warnings_SYS_4GAntenna                      Warnings = 124
	Warnings_SYS_SNUnset                        Warnings = 125
	Warnings_SYS_ParameterUnset                 Warnings = 126
	Warnings_SYS_ComGeneral                     Warnings = 127
	//*充电控制单元故障
	Warnings_CCM_EmergencyFault                    Warnings = 300
	Warnings_CCM_LockFault                         Warnings = 301
	Warnings_CCM_ACBreakerFault                    Warnings = 302
	Warnings_CCM_SPDFault                          Warnings = 303
	Warnings_CCM_SmogFault                         Warnings = 304
	Warnings_CCM_DoorFault                         Warnings = 305
	Warnings_CCM_WaterLogging                      Warnings = 306
	Warnings_CCM_DCContactorFault                  Warnings = 307
	Warnings_CCM_DCFuseFault                       Warnings = 308
	Warnings_CCM_InsulationFault                   Warnings = 309
	Warnings_CCM_InsulationWarning                 Warnings = 310
	Warnings_CCM_FanFault                          Warnings = 311
	Warnings_CCM_CC1Fault                          Warnings = 312
	Warnings_CCM_OverTemp                          Warnings = 313
	Warnings_CCM_ConnectorOverTemp                 Warnings = 314
	Warnings_CCM_GunHolderFault                    Warnings = 315
	Warnings_CCM_InputUnderVol                     Warnings = 316
	Warnings_CCM_InputOverVol                      Warnings = 317
	Warnings_CCM_AssisPowerFault                   Warnings = 318
	Warnings_CCM_BleedOffOverTime                  Warnings = 319
	Warnings_CCM_SwitchFault                       Warnings = 320
	Warnings_CCM_DCOuputUnderVol                   Warnings = 321
	Warnings_CCM_DCOutputOverVol                   Warnings = 322
	Warnings_CCM_DCOutputOverCur                   Warnings = 323
	Warnings_CCM_OverCurrent                       Warnings = 324
	Warnings_CCM_ChargingPlugOut                   Warnings = 325
	Warnings_CCM_ChargingFull                      Warnings = 326
	Warnings_CCM_CommunicationFault                Warnings = 327
	Warnings_CCM_PEFault                           Warnings = 328
	Warnings_CCM_TempSensorException               Warnings = 329
	Warnings_CCM_PDCircuitException                Warnings = 330
	Warnings_CCM_InputCircuitException             Warnings = 331
	Warnings_CCM_MainOutputCircuitException        Warnings = 332
	Warnings_CCM_AuxiliaryCircuitException         Warnings = 333
	Warnings_CCM_OutputContactAdhesion             Warnings = 334
	Warnings_CCM_OutputContactMisoperation         Warnings = 335
	Warnings_CCM_InputContactAdhesion              Warnings = 336
	Warnings_CCM_InputContactMisoperation          Warnings = 337
	Warnings_CCM_ParallelContactAdhesion           Warnings = 338
	Warnings_CCM_ParallelContactMisoperation       Warnings = 339
	Warnings_CCM_CANCommunicationFault             Warnings = 341
	Warnings_CCM_CAN1BusError                      Warnings = 342
	Warnings_CCM_CAN3CommunicationFault            Warnings = 343
	Warnings_CCM_SYSAuxiliaryPowerFailure          Warnings = 344
	Warnings_CCM_AuxiliaryPowerFault               Warnings = 345
	Warnings_CCM_ChargerFault                      Warnings = 346
	Warnings_CCM_TCUCommunicationFault             Warnings = 347
	Warnings_CCM_PowerControllerCommunicationFault Warnings = 348
	Warnings_CCM_ACFuseFault                       Warnings = 349
	Warnings_CCM_CPFault                           Warnings = 350
	Warnings_CCM_RCDSelfTest                       Warnings = 352
	Warnings_CCM_HighTemperature                   Warnings = 353
	Warnings_CCM_LowTemperature                    Warnings = 354
	Warnings_CCM_CPBroken                          Warnings = 355
	Warnings_CCM_Pwm                               Warnings = 356
	Warnings_CCM_OneLevelCurrentOverload           Warnings = 357
	Warnings_CCM_TwoLevelCurrentOverload           Warnings = 358
	Warnings_CCM_RelaySelfTest                     Warnings = 359
	Warnings_CCM_RelayStatus                       Warnings = 360
	Warnings_CCM_RelayControl                      Warnings = 361
	Warnings_CCM_WatchDocMCUDeadlock               Warnings = 362
	Warnings_CCM_SupervisorKeyModuleDeadlock       Warnings = 363
	Warnings_CCM_SupervisorUnKeyModuleDeadlock     Warnings = 364
	Warnings_CCM_DipValue                          Warnings = 365
	Warnings_CCM_PmicWatchDogMessage               Warnings = 366
	Warnings_CCM_RFIDModule                        Warnings = 367
	Warnings_CCM_BLEModule                         Warnings = 368
	Warnings_CCM_ShortCircuit                      Warnings = 369
	Warnings_CCM_CP12VMonitor                      Warnings = 373
	Warnings_CCM_ShortCircuitCharging              Warnings = 374
	Warnings_CCM_TemperatureSensorOpenCircuit      Warnings = 375
	Warnings_CCM_TemperatureSensorDriftCircuit     Warnings = 376
	Warnings_CCM_MultipleTemperatureOverload       Warnings = 377
	Warnings_CCM_OneLevelCurrentOverloadTwice      Warnings = 378
	Warnings_CCM_RelayPreDriverHighTemperature     Warnings = 379
	Warnings_CCM_RelayPreDriverOverload            Warnings = 380
	Warnings_CCM_L2OneLevelVoltageOverload         Warnings = 381
	Warnings_CCM_L2TwoLevelVoltageOverload         Warnings = 382
	Warnings_CCM_L2LongTimeVoltageOverload         Warnings = 383
	Warnings_CCM_L2VoltageOwe                      Warnings = 384
	Warnings_CCM_L2VoltageOweReduceFrequency       Warnings = 385
	Warnings_CCM_L2LongTimeVoltageOwe              Warnings = 386
	Warnings_CCM_L2OneLevelCurrentOverload         Warnings = 387
	Warnings_CCM_L2OneLevelCurrentOverloadTwice    Warnings = 388
	Warnings_CCM_L2TwoLevelCurrentOverload         Warnings = 389
	Warnings_CCM_L3OneLevelVoltageOverload         Warnings = 390
	Warnings_CCM_L3TwoLevelVoltageOverload         Warnings = 391
	Warnings_CCM_L3LongTimeVoltageOverload         Warnings = 392
	Warnings_CCM_L3VoltageOwe                      Warnings = 393
	Warnings_CCM_L3VoltageOweReduceFrequency       Warnings = 394
	Warnings_CCM_L3LongTimeVoltageOwe              Warnings = 395
	Warnings_CCM_L3OneLevelCurrentOverload         Warnings = 396
	Warnings_CCM_L3OneLevelCurrentOverloadTwice    Warnings = 397
	Warnings_CCM_L3TwoLevelCurrentOverload         Warnings = 398
	//*电动汽车BMS故障
	Warnings_BMS_CellOverVol               Warnings = 400
	Warnings_BMS_CellOverTemp              Warnings = 401
	Warnings_BMS_CommunicationFault        Warnings = 402
	Warnings_BMS_CommunicationOverTime     Warnings = 403
	Warnings_BMS_BatteryInversed           Warnings = 404
	Warnings_BMS_BatteryMissed             Warnings = 405
	Warnings_BMS_BatteryOverVol            Warnings = 406
	Warnings_BMS_OverSOC                   Warnings = 407
	Warnings_BMS_InsulationFault           Warnings = 408
	Warnings_BMS_ContactorFault            Warnings = 409
	Warnings_BMS_OverExternalVoltage       Warnings = 410
	Warnings_BMS_BatteryUnderVol           Warnings = 411
	Warnings_BMS_BatteryVoltageException   Warnings = 412
	Warnings_BMS_CellVoltageDiffFault      Warnings = 413
	Warnings_BMS_BusError                  Warnings = 416
	Warnings_BMS_CellSamplingPointOverTemp Warnings = 417
	//*电源模块故障
	Warnings_PSM_Fault                                 Warnings = 1200
	Warnings_PSM_ACInputFault                          Warnings = 1201
	Warnings_PSM_InputOverVol                          Warnings = 1202
	Warnings_PSM_InputUnderVol                         Warnings = 1203
	Warnings_PSM_InputLackingPhase                     Warnings = 1204
	Warnings_PSM_OutputShort                           Warnings = 1205
	Warnings_PSM_OutputOverCur                         Warnings = 1206
	Warnings_PSM_OutputOverVol                         Warnings = 1207
	Warnings_PSM_OutputUnderVol                        Warnings = 1208
	Warnings_PSM_OutputFault                           Warnings = 1209
	Warnings_PSM_OverTemp                              Warnings = 1210
	Warnings_PSM_CommunicationFault                    Warnings = 1211
	Warnings_PSM_FanFault                              Warnings = 1212
	Warnings_PSM_CloseException                        Warnings = 1213
	Warnings_PSM_RectifierNO0CommunicationFault        Warnings = 1214
	Warnings_PSM_RectifierNO1CommunicationFault        Warnings = 1215
	Warnings_PSM_RectifierNO2CommunicationFault        Warnings = 1216
	Warnings_PSM_RectifierNO3CommunicationFault        Warnings = 1217
	Warnings_PSM_RectifierNO4CommunicationFault        Warnings = 1218
	Warnings_PSM_RectifierNO5CommunicationFault        Warnings = 1219
	Warnings_PSM_RectifierNO6CommunicationFault        Warnings = 1220
	Warnings_PSM_RectifierNO7CommunicationFault        Warnings = 1221
	Warnings_PSM_RectifierNO8CommunicationFault        Warnings = 1222
	Warnings_PSM_RectifierNO9CommunicationFault        Warnings = 1223
	Warnings_PSM_RectifierNO10CommunicationFault       Warnings = 1224
	Warnings_PSM_RectifierNO11CommunicationFault       Warnings = 1225
	Warnings_PSM_RectifierNO12CommunicationFault       Warnings = 1226
	Warnings_PSM_RectifierNO0Fault                     Warnings = 1227
	Warnings_PSM_RectifierNO1Fault                     Warnings = 1228
	Warnings_PSM_RectifierNO2Fault                     Warnings = 1229
	Warnings_PSM_RectifierNO3Fault                     Warnings = 1230
	Warnings_PSM_RectifierNO4Fault                     Warnings = 1231
	Warnings_PSM_RectifierNO5Fault                     Warnings = 1232
	Warnings_PSM_RectifierNO6Fault                     Warnings = 1233
	Warnings_PSM_RectifierNO7Fault                     Warnings = 1234
	Warnings_PSM_RectifierNO8Fault                     Warnings = 1235
	Warnings_PSM_RectifierNO9Fault                     Warnings = 1236
	Warnings_PSM_RectifierNO10Fault                    Warnings = 1237
	Warnings_PSM_RectifierNO11Fault                    Warnings = 1238
	Warnings_PSM_RectifierNO12Fault                    Warnings = 1239
	Warnings_PSM_InsulationCommunicationFault          Warnings = 1240
	Warnings_PSM_PowerScreenCommunicationFault         Warnings = 1241
	Warnings_PSM_PowerIOFault                          Warnings = 1242
	Warnings_PSM_ACInputOverFrequency                  Warnings = 1243
	Warnings_PSM_ACInputUnderFrequency                 Warnings = 1244
	Warnings_PSM_TypeDisaccord                         Warnings = 1245
	Warnings_PSM_GroupChargingModuleOverTemp           Warnings = 1246
	Warnings_PSM_GroupChargingModuleCommunicationFault Warnings = 1247
	Warnings_PSM_Initialization                        Warnings = 1248
	Warnings_PSM_OutageFault                           Warnings = 1250
	Warnings_PSM_Inversed                              Warnings = 1251
	Warnings_PSM_InputPhaseImbalance                   Warnings = 1252
	Warnings_PSM_ProtectionWarning                     Warnings = 1253
	Warnings_PSM_ProtectionFault                       Warnings = 1254
	Warnings_PSM_PowerLimit                            Warnings = 1255
	Warnings_PSM_UnevenFlow                            Warnings = 1256
	Warnings_PSM_CurrentLimitWarning                   Warnings = 1257
	Warnings_PSM_CurrentLimitFault                     Warnings = 1258
	Warnings_PSM_OneLevelVoltageOverload               Warnings = 1259
	Warnings_PSM_OneLevelOweVoltage                    Warnings = 1260
	Warnings_PSM_TwoLevelOweVoltage                    Warnings = 1261
	Warnings_PSM_LongTimeVoltageOverload               Warnings = 1262
	Warnings_PSM_LongTimeVoltageOwe                    Warnings = 1263
	Warnings_PSM_ElectricFrequency                     Warnings = 1264
	Warnings_PSM_TwoLevelVoltageOverload               Warnings = 1265
	Warnings_PSM_VoltageOweReduceFrequency             Warnings = 1266
	Warnings_PSM_VoltageOwe                            Warnings = 1277
	Warnings_PSM_PMICWatchdogMessage                   Warnings = 1278
	Warnings_PSM_MultipleVoltageException              Warnings = 1279
	Warnings_PSM_RCD_AC_LEAK                           Warnings = 1280
	Warnings_PSM_RCD_DC_LEAK                           Warnings = 1281
)

var Warnings_name = map[int32]string{
	0:    "Other",
	100:  "SYS_StorageFull",
	101:  "SYS_4GFault",
	102:  "SYS_SimFault",
	103:  "SYS_DialFault",
	104:  "SYS_ConnectError",
	105:  "SYS_Disconnect",
	106:  "SYS_PowerFailure",
	107:  "SYS_MeterError",
	108:  "SYS_CardReaderError",
	109:  "SYS_LCDError",
	110:  "SYS_MeterMeasureException",
	111:  "SYS_MeterCollectsException",
	112:  "SYS_SoftFault",
	113:  "SYS_AirInletOverTemp",
	114:  "SYS_AirInletUnderTemp",
	115:  "SYS_AirOutletOverTemp",
	116:  "SYS_SlaveFault",
	117:  "SYS_StorageFailure",
	118:  "SYS_PersistenceMemoryInconsistency",
	119:  "SYS_ADCInit",
	120:  "SYS_ConnectCPO",
	121:  "SYS_Register",
	122:  "SYS_ComBoardOffline",
	123:  "SYS_GPSAntenna",
	124:  "SYS_4GAntenna",
	125:  "SYS_SNUnset",
	126:  "SYS_ParameterUnset",
	127:  "SYS_ComGeneral",
	300:  "CCM_EmergencyFault",
	301:  "CCM_LockFault",
	302:  "CCM_ACBreakerFault",
	303:  "CCM_SPDFault",
	304:  "CCM_SmogFault",
	305:  "CCM_DoorFault",
	306:  "CCM_WaterLogging",
	307:  "CCM_DCContactorFault",
	308:  "CCM_DCFuseFault",
	309:  "CCM_InsulationFault",
	310:  "CCM_InsulationWarning",
	311:  "CCM_FanFault",
	312:  "CCM_CC1Fault",
	313:  "CCM_OverTemp",
	314:  "CCM_ConnectorOverTemp",
	315:  "CCM_GunHolderFault",
	316:  "CCM_InputUnderVol",
	317:  "CCM_InputOverVol",
	318:  "CCM_AssisPowerFault",
	319:  "CCM_BleedOffOverTime",
	320:  "CCM_SwitchFault",
	321:  "CCM_DCOuputUnderVol",
	322:  "CCM_DCOutputOverVol",
	323:  "CCM_DCOutputOverCur",
	324:  "CCM_OverCurrent",
	325:  "CCM_ChargingPlugOut",
	326:  "CCM_ChargingFull",
	327:  "CCM_CommunicationFault",
	328:  "CCM_PEFault",
	329:  "CCM_TempSensorException",
	330:  "CCM_PDCircuitException",
	331:  "CCM_InputCircuitException",
	332:  "CCM_MainOutputCircuitException",
	333:  "CCM_AuxiliaryCircuitException",
	334:  "CCM_OutputContactAdhesion",
	335:  "CCM_OutputContactMisoperation",
	336:  "CCM_InputContactAdhesion",
	337:  "CCM_InputContactMisoperation",
	338:  "CCM_ParallelContactAdhesion",
	339:  "CCM_ParallelContactMisoperation",
	341:  "CCM_CANCommunicationFault",
	342:  "CCM_CAN1BusError",
	343:  "CCM_CAN3CommunicationFault",
	344:  "CCM_SYSAuxiliaryPowerFailure",
	345:  "CCM_AuxiliaryPowerFault",
	346:  "CCM_ChargerFault",
	347:  "CCM_TCUCommunicationFault",
	348:  "CCM_PowerControllerCommunicationFault",
	349:  "CCM_ACFuseFault",
	350:  "CCM_CPFault",
	352:  "CCM_RCDSelfTest",
	353:  "CCM_HighTemperature",
	354:  "CCM_LowTemperature",
	355:  "CCM_CPBroken",
	356:  "CCM_Pwm",
	357:  "CCM_OneLevelCurrentOverload",
	358:  "CCM_TwoLevelCurrentOverload",
	359:  "CCM_RelaySelfTest",
	360:  "CCM_RelayStatus",
	361:  "CCM_RelayControl",
	362:  "CCM_WatchDocMCUDeadlock",
	363:  "CCM_SupervisorKeyModuleDeadlock",
	364:  "CCM_SupervisorUnKeyModuleDeadlock",
	365:  "CCM_DipValue",
	366:  "CCM_PmicWatchDogMessage",
	367:  "CCM_RFIDModule",
	368:  "CCM_BLEModule",
	369:  "CCM_ShortCircuit",
	373:  "CCM_CP12VMonitor",
	374:  "CCM_ShortCircuitCharging",
	375:  "CCM_TemperatureSensorOpenCircuit",
	376:  "CCM_TemperatureSensorDriftCircuit",
	377:  "CCM_MultipleTemperatureOverload",
	378:  "CCM_OneLevelCurrentOverloadTwice",
	379:  "CCM_RelayPreDriverHighTemperature",
	380:  "CCM_RelayPreDriverOverload",
	381:  "CCM_L2OneLevelVoltageOverload",
	382:  "CCM_L2TwoLevelVoltageOverload",
	383:  "CCM_L2LongTimeVoltageOverload",
	384:  "CCM_L2VoltageOwe",
	385:  "CCM_L2VoltageOweReduceFrequency",
	386:  "CCM_L2LongTimeVoltageOwe",
	387:  "CCM_L2OneLevelCurrentOverload",
	388:  "CCM_L2OneLevelCurrentOverloadTwice",
	389:  "CCM_L2TwoLevelCurrentOverload",
	390:  "CCM_L3OneLevelVoltageOverload",
	391:  "CCM_L3TwoLevelVoltageOverload",
	392:  "CCM_L3LongTimeVoltageOverload",
	393:  "CCM_L3VoltageOwe",
	394:  "CCM_L3VoltageOweReduceFrequency",
	395:  "CCM_L3LongTimeVoltageOwe",
	396:  "CCM_L3OneLevelCurrentOverload",
	397:  "CCM_L3OneLevelCurrentOverloadTwice",
	398:  "CCM_L3TwoLevelCurrentOverload",
	400:  "BMS_CellOverVol",
	401:  "BMS_CellOverTemp",
	402:  "BMS_CommunicationFault",
	403:  "BMS_CommunicationOverTime",
	404:  "BMS_BatteryInversed",
	405:  "BMS_BatteryMissed",
	406:  "BMS_BatteryOverVol",
	407:  "BMS_OverSOC",
	408:  "BMS_InsulationFault",
	409:  "BMS_ContactorFault",
	410:  "BMS_OverExternalVoltage",
	411:  "BMS_BatteryUnderVol",
	412:  "BMS_BatteryVoltageException",
	413:  "BMS_CellVoltageDiffFault",
	416:  "BMS_BusError",
	417:  "BMS_CellSamplingPointOverTemp",
	1200: "PSM_Fault",
	1201: "PSM_ACInputFault",
	1202: "PSM_InputOverVol",
	1203: "PSM_InputUnderVol",
	1204: "PSM_InputLackingPhase",
	1205: "PSM_OutputShort",
	1206: "PSM_OutputOverCur",
	1207: "PSM_OutputOverVol",
	1208: "PSM_OutputUnderVol",
	1209: "PSM_OutputFault",
	1210: "PSM_OverTemp",
	1211: "PSM_CommunicationFault",
	1212: "PSM_FanFault",
	1213: "PSM_CloseException",
	1214: "PSM_RectifierNO0CommunicationFault",
	1215: "PSM_RectifierNO1CommunicationFault",
	1216: "PSM_RectifierNO2CommunicationFault",
	1217: "PSM_RectifierNO3CommunicationFault",
	1218: "PSM_RectifierNO4CommunicationFault",
	1219: "PSM_RectifierNO5CommunicationFault",
	1220: "PSM_RectifierNO6CommunicationFault",
	1221: "PSM_RectifierNO7CommunicationFault",
	1222: "PSM_RectifierNO8CommunicationFault",
	1223: "PSM_RectifierNO9CommunicationFault",
	1224: "PSM_RectifierNO10CommunicationFault",
	1225: "PSM_RectifierNO11CommunicationFault",
	1226: "PSM_RectifierNO12CommunicationFault",
	1227: "PSM_RectifierNO0Fault",
	1228: "PSM_RectifierNO1Fault",
	1229: "PSM_RectifierNO2Fault",
	1230: "PSM_RectifierNO3Fault",
	1231: "PSM_RectifierNO4Fault",
	1232: "PSM_RectifierNO5Fault",
	1233: "PSM_RectifierNO6Fault",
	1234: "PSM_RectifierNO7Fault",
	1235: "PSM_RectifierNO8Fault",
	1236: "PSM_RectifierNO9Fault",
	1237: "PSM_RectifierNO10Fault",
	1238: "PSM_RectifierNO11Fault",
	1239: "PSM_RectifierNO12Fault",
	1240: "PSM_InsulationCommunicationFault",
	1241: "PSM_PowerScreenCommunicationFault",
	1242: "PSM_PowerIOFault",
	1243: "PSM_ACInputOverFrequency",
	1244: "PSM_ACInputUnderFrequency",
	1245: "PSM_TypeDisaccord",
	1246: "PSM_GroupChargingModuleOverTemp",
	1247: "PSM_GroupChargingModuleCommunicationFault",
	1248: "PSM_Initialization",
	1250: "PSM_OutageFault",
	1251: "PSM_Inversed",
	1252: "PSM_InputPhaseImbalance",
	1253: "PSM_ProtectionWarning",
	1254: "PSM_ProtectionFault",
	1255: "PSM_PowerLimit",
	1256: "PSM_UnevenFlow",
	1257: "PSM_CurrentLimitWarning",
	1258: "PSM_CurrentLimitFault",
	1259: "PSM_OneLevelVoltageOverload",
	1260: "PSM_OneLevelOweVoltage",
	1261: "PSM_TwoLevelOweVoltage",
	1262: "PSM_LongTimeVoltageOverload",
	1263: "PSM_LongTimeVoltageOwe",
	1264: "PSM_ElectricFrequency",
	1265: "PSM_TwoLevelVoltageOverload",
	1266: "PSM_VoltageOweReduceFrequency",
	1277: "PSM_VoltageOwe",
	1278: "PSM_PMICWatchdogMessage",
	1279: "PSM_MultipleVoltageException",
	1280: "PSM_RCD_AC_LEAK",
	1281: "PSM_RCD_DC_LEAK",
}

var Warnings_value = map[string]int32{
	"Other":                                     0,
	"SYS_StorageFull":                           100,
	"SYS_4GFault":                               101,
	"SYS_SimFault":                              102,
	"SYS_DialFault":                             103,
	"SYS_ConnectError":                          104,
	"SYS_Disconnect":                            105,
	"SYS_PowerFailure":                          106,
	"SYS_MeterError":                            107,
	"SYS_CardReaderError":                       108,
	"SYS_LCDError":                              109,
	"SYS_MeterMeasureException":                 110,
	"SYS_MeterCollectsException":                111,
	"SYS_SoftFault":                             112,
	"SYS_AirInletOverTemp":                      113,
	"SYS_AirInletUnderTemp":                     114,
	"SYS_AirOutletOverTemp":                     115,
	"SYS_SlaveFault":                            116,
	"SYS_StorageFailure":                        117,
	"SYS_PersistenceMemoryInconsistency":        118,
	"SYS_ADCInit":                               119,
	"SYS_ConnectCPO":                            120,
	"SYS_Register":                              121,
	"SYS_ComBoardOffline":                       122,
	"SYS_GPSAntenna":                            123,
	"SYS_4GAntenna":                             124,
	"SYS_SNUnset":                               125,
	"SYS_ParameterUnset":                        126,
	"SYS_ComGeneral":                            127,
	"CCM_EmergencyFault":                        300,
	"CCM_LockFault":                             301,
	"CCM_ACBreakerFault":                        302,
	"CCM_SPDFault":                              303,
	"CCM_SmogFault":                             304,
	"CCM_DoorFault":                             305,
	"CCM_WaterLogging":                          306,
	"CCM_DCContactorFault":                      307,
	"CCM_DCFuseFault":                           308,
	"CCM_InsulationFault":                       309,
	"CCM_InsulationWarning":                     310,
	"CCM_FanFault":                              311,
	"CCM_CC1Fault":                              312,
	"CCM_OverTemp":                              313,
	"CCM_ConnectorOverTemp":                     314,
	"CCM_GunHolderFault":                        315,
	"CCM_InputUnderVol":                         316,
	"CCM_InputOverVol":                          317,
	"CCM_AssisPowerFault":                       318,
	"CCM_BleedOffOverTime":                      319,
	"CCM_SwitchFault":                           320,
	"CCM_DCOuputUnderVol":                       321,
	"CCM_DCOutputOverVol":                       322,
	"CCM_DCOutputOverCur":                       323,
	"CCM_OverCurrent":                           324,
	"CCM_ChargingPlugOut":                       325,
	"CCM_ChargingFull":                          326,
	"CCM_CommunicationFault":                    327,
	"CCM_PEFault":                               328,
	"CCM_TempSensorException":                   329,
	"CCM_PDCircuitException":                    330,
	"CCM_InputCircuitException":                 331,
	"CCM_MainOutputCircuitException":            332,
	"CCM_AuxiliaryCircuitException":             333,
	"CCM_OutputContactAdhesion":                 334,
	"CCM_OutputContactMisoperation":             335,
	"CCM_InputContactAdhesion":                  336,
	"CCM_InputContactMisoperation":              337,
	"CCM_ParallelContactAdhesion":               338,
	"CCM_ParallelContactMisoperation":           339,
	"CCM_CANCommunicationFault":                 341,
	"CCM_CAN1BusError":                          342,
	"CCM_CAN3CommunicationFault":                343,
	"CCM_SYSAuxiliaryPowerFailure":              344,
	"CCM_AuxiliaryPowerFault":                   345,
	"CCM_ChargerFault":                          346,
	"CCM_TCUCommunicationFault":                 347,
	"CCM_PowerControllerCommunicationFault":     348,
	"CCM_ACFuseFault":                           349,
	"CCM_CPFault":                               350,
	"CCM_RCDSelfTest":                           352,
	"CCM_HighTemperature":                       353,
	"CCM_LowTemperature":                        354,
	"CCM_CPBroken":                              355,
	"CCM_Pwm":                                   356,
	"CCM_OneLevelCurrentOverload":               357,
	"CCM_TwoLevelCurrentOverload":               358,
	"CCM_RelaySelfTest":                         359,
	"CCM_RelayStatus":                           360,
	"CCM_RelayControl":                          361,
	"CCM_WatchDocMCUDeadlock":                   362,
	"CCM_SupervisorKeyModuleDeadlock":           363,
	"CCM_SupervisorUnKeyModuleDeadlock":         364,
	"CCM_DipValue":                              365,
	"CCM_PmicWatchDogMessage":                   366,
	"CCM_RFIDModule":                            367,
	"CCM_BLEModule":                             368,
	"CCM_ShortCircuit":                          369,
	"CCM_CP12VMonitor":                          373,
	"CCM_ShortCircuitCharging":                  374,
	"CCM_TemperatureSensorOpenCircuit":          375,
	"CCM_TemperatureSensorDriftCircuit":         376,
	"CCM_MultipleTemperatureOverload":           377,
	"CCM_OneLevelCurrentOverloadTwice":          378,
	"CCM_RelayPreDriverHighTemperature":         379,
	"CCM_RelayPreDriverOverload":                380,
	"CCM_L2OneLevelVoltageOverload":             381,
	"CCM_L2TwoLevelVoltageOverload":             382,
	"CCM_L2LongTimeVoltageOverload":             383,
	"CCM_L2VoltageOwe":                          384,
	"CCM_L2VoltageOweReduceFrequency":           385,
	"CCM_L2LongTimeVoltageOwe":                  386,
	"CCM_L2OneLevelCurrentOverload":             387,
	"CCM_L2OneLevelCurrentOverloadTwice":        388,
	"CCM_L2TwoLevelCurrentOverload":             389,
	"CCM_L3OneLevelVoltageOverload":             390,
	"CCM_L3TwoLevelVoltageOverload":             391,
	"CCM_L3LongTimeVoltageOverload":             392,
	"CCM_L3VoltageOwe":                          393,
	"CCM_L3VoltageOweReduceFrequency":           394,
	"CCM_L3LongTimeVoltageOwe":                  395,
	"CCM_L3OneLevelCurrentOverload":             396,
	"CCM_L3OneLevelCurrentOverloadTwice":        397,
	"CCM_L3TwoLevelCurrentOverload":             398,
	"BMS_CellOverVol":                           400,
	"BMS_CellOverTemp":                          401,
	"BMS_CommunicationFault":                    402,
	"BMS_CommunicationOverTime":                 403,
	"BMS_BatteryInversed":                       404,
	"BMS_BatteryMissed":                         405,
	"BMS_BatteryOverVol":                        406,
	"BMS_OverSOC":                               407,
	"BMS_InsulationFault":                       408,
	"BMS_ContactorFault":                        409,
	"BMS_OverExternalVoltage":                   410,
	"BMS_BatteryUnderVol":                       411,
	"BMS_BatteryVoltageException":               412,
	"BMS_CellVoltageDiffFault":                  413,
	"BMS_BusError":                              416,
	"BMS_CellSamplingPointOverTemp":             417,
	"PSM_Fault":                                 1200,
	"PSM_ACInputFault":                          1201,
	"PSM_InputOverVol":                          1202,
	"PSM_InputUnderVol":                         1203,
	"PSM_InputLackingPhase":                     1204,
	"PSM_OutputShort":                           1205,
	"PSM_OutputOverCur":                         1206,
	"PSM_OutputOverVol":                         1207,
	"PSM_OutputUnderVol":                        1208,
	"PSM_OutputFault":                           1209,
	"PSM_OverTemp":                              1210,
	"PSM_CommunicationFault":                    1211,
	"PSM_FanFault":                              1212,
	"PSM_CloseException":                        1213,
	"PSM_RectifierNO0CommunicationFault":        1214,
	"PSM_RectifierNO1CommunicationFault":        1215,
	"PSM_RectifierNO2CommunicationFault":        1216,
	"PSM_RectifierNO3CommunicationFault":        1217,
	"PSM_RectifierNO4CommunicationFault":        1218,
	"PSM_RectifierNO5CommunicationFault":        1219,
	"PSM_RectifierNO6CommunicationFault":        1220,
	"PSM_RectifierNO7CommunicationFault":        1221,
	"PSM_RectifierNO8CommunicationFault":        1222,
	"PSM_RectifierNO9CommunicationFault":        1223,
	"PSM_RectifierNO10CommunicationFault":       1224,
	"PSM_RectifierNO11CommunicationFault":       1225,
	"PSM_RectifierNO12CommunicationFault":       1226,
	"PSM_RectifierNO0Fault":                     1227,
	"PSM_RectifierNO1Fault":                     1228,
	"PSM_RectifierNO2Fault":                     1229,
	"PSM_RectifierNO3Fault":                     1230,
	"PSM_RectifierNO4Fault":                     1231,
	"PSM_RectifierNO5Fault":                     1232,
	"PSM_RectifierNO6Fault":                     1233,
	"PSM_RectifierNO7Fault":                     1234,
	"PSM_RectifierNO8Fault":                     1235,
	"PSM_RectifierNO9Fault":                     1236,
	"PSM_RectifierNO10Fault":                    1237,
	"PSM_RectifierNO11Fault":                    1238,
	"PSM_RectifierNO12Fault":                    1239,
	"PSM_InsulationCommunicationFault":          1240,
	"PSM_PowerScreenCommunicationFault":         1241,
	"PSM_PowerIOFault":                          1242,
	"PSM_ACInputOverFrequency":                  1243,
	"PSM_ACInputUnderFrequency":                 1244,
	"PSM_TypeDisaccord":                         1245,
	"PSM_GroupChargingModuleOverTemp":           1246,
	"PSM_GroupChargingModuleCommunicationFault": 1247,
	"PSM_Initialization":                        1248,
	"PSM_OutageFault":                           1250,
	"PSM_Inversed":                              1251,
	"PSM_InputPhaseImbalance":                   1252,
	"PSM_ProtectionWarning":                     1253,
	"PSM_ProtectionFault":                       1254,
	"PSM_PowerLimit":                            1255,
	"PSM_UnevenFlow":                            1256,
	"PSM_CurrentLimitWarning":                   1257,
	"PSM_CurrentLimitFault":                     1258,
	"PSM_OneLevelVoltageOverload":               1259,
	"PSM_OneLevelOweVoltage":                    1260,
	"PSM_TwoLevelOweVoltage":                    1261,
	"PSM_LongTimeVoltageOverload":               1262,
	"PSM_LongTimeVoltageOwe":                    1263,
	"PSM_ElectricFrequency":                     1264,
	"PSM_TwoLevelVoltageOverload":               1265,
	"PSM_VoltageOweReduceFrequency":             1266,
	"PSM_VoltageOwe":                            1277,
	"PSM_PMICWatchdogMessage":                   1278,
	"PSM_MultipleVoltageException":              1279,
	"PSM_RCD_AC_LEAK":                           1280,
	"PSM_RCD_DC_LEAK":                           1281,
}

func (x Warnings) String() string {
	return proto.EnumName(Warnings_name, int32(x))
}

func (Warnings) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f1bd32ed1d0030e8, []int{0}
}

func init() {
	proto.RegisterEnum("Goiot.Protocol.Charger.Warning.Warnings", Warnings_name, Warnings_value)
}

func init() { proto.RegisterFile("charger/warning.proto", fileDescriptor_f1bd32ed1d0030e8) }

var fileDescriptor_f1bd32ed1d0030e8 = []byte{
	// 2261 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x58, 0x57, 0x70, 0x1c, 0xc7,
	0xd1, 0xfe, 0x6b, 0xaf, 0x7e, 0xdb, 0x3b, 0xa2, 0xc8, 0xe1, 0x30, 0x53, 0x22, 0x25, 0x51, 0x96,
	0x6c, 0xeb, 0x01, 0x30, 0x09, 0x2a, 0x3d, 0x1e, 0xf6, 0x00, 0x10, 0xa5, 0x3b, 0xde, 0x15, 0x16,
	0x20, 0x4b, 0x7e, 0x61, 0x8d, 0xf6, 0x06, 0x87, 0x11, 0xf6, 0x76, 0x8e, 0xb3, 0xbb, 0x38, 0x42,
	0x4e, 0x92, 0x73, 0x2e, 0xe7, 0xec, 0x2a, 0xfb, 0xc9, 0x2f, 0x8e, 0xca, 0xb2, 0x72, 0xb4, 0x72,
	0xb4, 0x15, 0xac, 0x60, 0x59, 0x39, 0x07, 0x57, 0x39, 0x5b, 0x92, 0x6b, 0x76, 0x76, 0x7a, 0xf7,
	0x0e, 0x03, 0xbd, 0x01, 0xdf, 0xf7, 0x4d, 0xf7, 0x74, 0xf7, 0x6c, 0x4f, 0xcf, 0xa1, 0x4d, 0xc1,
	0x02, 0x95, 0x1d, 0x26, 0x47, 0xfb, 0x54, 0x46, 0x3c, 0xea, 0x8c, 0xf4, 0xa4, 0x48, 0x04, 0xd9,
	0x39, 0x25, 0xb8, 0x48, 0x46, 0x5a, 0xea, 0x9f, 0x40, 0x84, 0x23, 0x9e, 0x56, 0x8d, 0x1c, 0xd4,
	0xaa, 0x53, 0x7e, 0x3e, 0x82, 0xde, 0x97, 0xff, 0x1d, 0x13, 0x17, 0xfd, 0x7f, 0x33, 0x59, 0x60,
	0x12, 0xff, 0x1f, 0xd9, 0x80, 0xd6, 0xf9, 0x67, 0xfb, 0x87, 0xfc, 0x44, 0x48, 0xda, 0x61, 0x93,
	0x69, 0x18, 0xe2, 0x36, 0x59, 0x87, 0x8e, 0x52, 0xe0, 0xde, 0xa9, 0x49, 0x9a, 0x86, 0x09, 0x66,
	0x04, 0xa3, 0x35, 0x99, 0x8a, 0x77, 0x35, 0x32, 0x4f, 0xd6, 0xa3, 0xa3, 0x15, 0x52, 0xe3, 0x34,
	0xd4, 0x50, 0x87, 0x6c, 0x44, 0x58, 0x41, 0x9e, 0x88, 0x22, 0x16, 0x24, 0x13, 0x52, 0x0a, 0x89,
	0x17, 0x08, 0x41, 0x6b, 0xb5, 0x30, 0x0e, 0x34, 0x81, 0xb9, 0x51, 0xb6, 0x44, 0x9f, 0xc9, 0x49,
	0xca, 0xc3, 0x54, 0x32, 0x7c, 0xae, 0x51, 0x36, 0x58, 0xc2, 0xa4, 0x5e, 0xbd, 0x48, 0xb6, 0xa0,
	0x0d, 0x99, 0x4d, 0x2a, 0xdb, 0x33, 0x8c, 0xb6, 0x0d, 0x11, 0x9a, 0x1d, 0xd5, 0xbd, 0x9a, 0x46,
	0xba, 0x64, 0x07, 0xda, 0x06, 0xcb, 0x1b, 0x8c, 0xc6, 0xa9, 0x64, 0x13, 0x47, 0x02, 0xd6, 0x4b,
	0xb8, 0x88, 0x70, 0x44, 0x76, 0xa2, 0xed, 0x40, 0x7b, 0x22, 0x0c, 0x59, 0x90, 0xc4, 0x05, 0x2f,
	0x4c, 0x40, 0xbe, 0x98, 0x4f, 0x74, 0x40, 0x3d, 0xb2, 0x15, 0x6d, 0x54, 0x50, 0x95, 0xcb, 0xe9,
	0x28, 0x64, 0x49, 0x73, 0x89, 0xc9, 0x59, 0xd6, 0xed, 0xe1, 0xc3, 0x64, 0x1b, 0xda, 0x54, 0x66,
	0xe6, 0xa2, 0x76, 0x4e, 0xc9, 0x12, 0xd5, 0x4c, 0x93, 0xf2, 0xaa, 0xd8, 0x04, 0xe8, 0x87, 0x74,
	0x89, 0x69, 0x1f, 0x09, 0xd9, 0x8c, 0x48, 0x39, 0xff, 0x79, 0x32, 0x52, 0x72, 0x32, 0xda, 0x95,
	0xa5, 0x88, 0xc9, 0x98, 0xc7, 0x09, 0x8b, 0x02, 0xd6, 0x60, 0x5d, 0x21, 0x97, 0xa7, 0xa3, 0x40,
	0x44, 0x39, 0xb4, 0x8c, 0x97, 0x4c, 0xa9, 0xaa, 0x35, 0x6f, 0x3a, 0xe2, 0x09, 0xee, 0x1b, 0x27,
	0x79, 0x15, 0xbc, 0x56, 0x13, 0x1f, 0x31, 0xc9, 0x9a, 0x61, 0x1d, 0xb5, 0x50, 0xe2, 0x65, 0xc8,
	0xab, 0xe8, 0x8e, 0x0b, 0x2a, 0xdb, 0xcd, 0xf9, 0xf9, 0x90, 0x47, 0x0c, 0x9f, 0x67, 0x96, 0x4f,
	0xb5, 0xfc, 0x6a, 0x94, 0xb0, 0x28, 0xa2, 0xf8, 0xa3, 0x26, 0x35, 0x7b, 0xa7, 0x0c, 0xf4, 0x31,
	0xe3, 0xd6, 0xdf, 0x3f, 0x17, 0xc5, 0x2c, 0xc1, 0x1f, 0x37, 0x71, 0xb4, 0xa8, 0xa4, 0x5d, 0x95,
	0x62, 0x8d, 0x7f, 0xa2, 0xd8, 0x4e, 0x77, 0x8a, 0x45, 0x4c, 0xd2, 0x10, 0x7f, 0x92, 0x6c, 0x41,
	0xc4, 0xf3, 0x1a, 0x87, 0x26, 0xba, 0x4c, 0x76, 0x54, 0x18, 0x3a, 0x17, 0xbf, 0x70, 0x08, 0x41,
	0x47, 0x2b, 0xa2, 0x2e, 0x82, 0x45, 0x8d, 0xfd, 0xd2, 0x31, 0xe2, 0xaa, 0x37, 0x2e, 0x19, 0x5d,
	0x54, 0xe7, 0x45, 0x11, 0xbf, 0x72, 0xc8, 0x7a, 0xb4, 0x46, 0x11, 0x7e, 0xab, 0xa6, 0xa1, 0x5f,
	0xc3, 0x7a, 0xbf, 0x2b, 0x3a, 0x1a, 0xfb, 0x0d, 0x60, 0x35, 0x21, 0xf2, 0xa5, 0xbf, 0x75, 0xc8,
	0x26, 0x84, 0x15, 0x76, 0x90, 0x26, 0x4c, 0xd6, 0x45, 0xa7, 0xc3, 0xa3, 0x0e, 0xbe, 0xd0, 0x21,
	0xdb, 0xd0, 0xc6, 0x4c, 0xea, 0x79, 0x22, 0x4a, 0x68, 0x90, 0x98, 0x15, 0x17, 0x39, 0x64, 0x23,
	0x5a, 0xa7, 0xa9, 0xc9, 0x34, 0xce, 0x6b, 0x77, 0xb1, 0x43, 0xb6, 0xa2, 0x0d, 0x0a, 0x9d, 0x8e,
	0xe2, 0x34, 0xa4, 0xea, 0x1c, 0x69, 0xe6, 0x12, 0x87, 0x6c, 0x47, 0x9b, 0x06, 0x99, 0xfc, 0xdb,
	0xc3, 0x97, 0xc2, 0xc6, 0x27, 0x69, 0x2e, 0xbf, 0x0c, 0x20, 0xcf, 0xdb, 0xad, 0xa1, 0xcb, 0x01,
	0x82, 0xe3, 0x73, 0x05, 0x18, 0xcd, 0x4b, 0x2b, 0x24, 0x70, 0xbf, 0x83, 0x34, 0x4d, 0xa5, 0xd1,
	0x3e, 0x11, 0xb6, 0x4d, 0x9a, 0xae, 0x74, 0xc8, 0x66, 0xb4, 0x5e, 0xef, 0xa4, 0x97, 0xea, 0x73,
	0x7a, 0x40, 0x84, 0xf8, 0x2a, 0xc8, 0x41, 0x86, 0x2b, 0x43, 0x0a, 0xbe, 0x1a, 0x42, 0xaa, 0xc6,
	0x31, 0x8f, 0xf3, 0xef, 0x53, 0x19, 0xba, 0x06, 0xb2, 0x33, 0x1e, 0x32, 0xa6, 0xce, 0x4b, 0xe6,
	0x9c, 0x77, 0x19, 0xbe, 0x16, 0xb2, 0xe3, 0xf7, 0x79, 0x12, 0x2c, 0xe8, 0x05, 0xd7, 0x81, 0xa9,
	0x9a, 0xd7, 0x4c, 0xcb, 0xbe, 0xaf, 0x1f, 0x60, 0x92, 0x92, 0xfb, 0x1b, 0xac, 0x8c, 0x97, 0x4a,
	0x7c, 0x23, 0xf8, 0xc8, 0x11, 0xc9, 0xa2, 0x04, 0xdf, 0x04, 0xfa, 0xac, 0xdb, 0xf1, 0xa8, 0xd3,
	0x0a, 0xd3, 0x4e, 0x33, 0x4d, 0xf0, 0xcd, 0x10, 0x9f, 0x61, 0xb2, 0xce, 0x76, 0x8b, 0x43, 0x8e,
	0x41, 0x9b, 0x75, 0x0e, 0xbb, 0xdd, 0x34, 0xe2, 0x41, 0xa9, 0x6a, 0xbf, 0x77, 0x08, 0x46, 0x47,
	0x29, 0xb2, 0x35, 0xa1, 0x91, 0x5b, 0x1d, 0x72, 0x2c, 0xda, 0xa2, 0x10, 0x95, 0x65, 0x9f, 0x45,
	0xb1, 0x90, 0x45, 0xcb, 0xb8, 0x0d, 0x8c, 0xb5, 0x6a, 0x1e, 0x97, 0x41, 0xca, 0x93, 0x82, 0xbc,
	0xdd, 0x21, 0x3b, 0xd1, 0x36, 0x48, 0xf0, 0x0a, 0xfe, 0x0e, 0x87, 0x9c, 0x88, 0x76, 0x2a, 0xbe,
	0x41, 0x79, 0xa4, 0x83, 0x5d, 0x21, 0xba, 0xd3, 0x21, 0xbb, 0xd0, 0x8e, 0xac, 0x1c, 0xe9, 0x11,
	0x1e, 0x72, 0x2a, 0x97, 0x57, 0x68, 0xee, 0x02, 0x47, 0xb9, 0x11, 0x7d, 0x74, 0xab, 0xed, 0x05,
	0x16, 0x2b, 0xfe, 0x6e, 0xb0, 0x31, 0xc0, 0x37, 0x78, 0x2c, 0x7a, 0x4c, 0x66, 0xe1, 0xe3, 0x7b,
	0x1c, 0xb2, 0x03, 0x6d, 0x2d, 0x36, 0x3b, 0x64, 0xe2, 0x5e, 0x87, 0x9c, 0x80, 0x8e, 0x1d, 0xa6,
	0x07, 0x2c, 0xdc, 0xe7, 0x90, 0xe3, 0xd1, 0x31, 0x59, 0x2e, 0xa8, 0xa4, 0x61, 0xc8, 0xc2, 0x61,
	0x23, 0xf7, 0x3b, 0xe4, 0xfd, 0xe8, 0x38, 0x8b, 0x62, 0xc0, 0xce, 0x03, 0x10, 0x8d, 0x57, 0xdd,
	0x6f, 0xa9, 0xd1, 0x1f, 0x8b, 0xba, 0x56, 0xf7, 0xef, 0x1e, 0x4f, 0x63, 0xdd, 0xfc, 0x1f, 0x74,
	0xc8, 0x71, 0x68, 0x7b, 0x0e, 0x8f, 0x59, 0xd6, 0x3d, 0x04, 0x21, 0xf8, 0x67, 0xfb, 0x90, 0xcc,
	0x81, 0xfb, 0xe7, 0x61, 0x28, 0xf6, 0x30, 0xaf, 0x0c, 0x3c, 0x32, 0x78, 0xa0, 0x0c, 0xfc, 0x27,
	0xd8, 0xef, 0xac, 0x37, 0x67, 0xf1, 0xfb, 0xa8, 0x43, 0x4e, 0x41, 0x27, 0x65, 0x51, 0x2b, 0x5b,
	0x2a, 0x64, 0xa9, 0xee, 0x1e, 0x69, 0xd1, 0x3e, 0x06, 0x67, 0xbc, 0x5a, 0xea, 0x32, 0x8f, 0xc3,
	0xa9, 0xf4, 0x5a, 0x1a, 0x79, 0x02, 0x74, 0x33, 0x5e, 0xcd, 0x67, 0xe1, 0xfc, 0x2c, 0x8b, 0x13,
	0xfc, 0x24, 0x7c, 0x0b, 0xfb, 0x78, 0x67, 0x41, 0x9d, 0x57, 0x95, 0x53, 0x15, 0xd8, 0x5f, 0xa0,
	0x39, 0xd4, 0x45, 0xbf, 0x4c, 0x3c, 0x55, 0xf4, 0x9d, 0xd6, 0xb8, 0x14, 0x8b, 0x2c, 0xc2, 0x4f,
	0x3b, 0x64, 0x0d, 0x7a, 0x6f, 0xb6, 0xdf, 0x7e, 0x17, 0x3f, 0x03, 0x55, 0x6d, 0x46, 0xac, 0xce,
	0x96, 0x58, 0x98, 0x7f, 0x79, 0xea, 0x23, 0x0c, 0x05, 0x6d, 0xe3, 0x67, 0x41, 0x31, 0xdb, 0x17,
	0x56, 0xc5, 0x73, 0xd0, 0x81, 0x66, 0x58, 0x48, 0x97, 0x61, 0xbf, 0xcf, 0x17, 0x51, 0x64, 0x78,
	0x42, 0x93, 0x34, 0xc6, 0x2f, 0x40, 0x9a, 0x33, 0x34, 0xcf, 0x17, 0x7e, 0x11, 0x6a, 0x73, 0x90,
	0x26, 0xc1, 0x42, 0x4d, 0x04, 0x0d, 0x6f, 0xae, 0xc6, 0x68, 0x3b, 0x14, 0xc1, 0x22, 0x7e, 0x09,
	0x8e, 0x96, 0x9f, 0xf6, 0x98, 0x5c, 0xe2, 0xb1, 0x90, 0x67, 0xb1, 0xe5, 0x86, 0x68, 0xa7, 0x21,
	0x03, 0xd5, 0xcb, 0x0e, 0x39, 0x19, 0x9d, 0x30, 0xa8, 0x9a, 0x8b, 0x56, 0xea, 0x5e, 0x81, 0xac,
	0xd4, 0x78, 0xef, 0x00, 0x0d, 0x53, 0x86, 0x5f, 0x05, 0xf7, 0xad, 0x2e, 0x0f, 0xf2, 0x2d, 0x74,
	0x1a, 0x2c, 0x8e, 0x69, 0x87, 0xe1, 0xd7, 0x1c, 0xb2, 0x01, 0xad, 0xcd, 0xf6, 0x3c, 0x39, 0x5d,
	0xd3, 0xd6, 0xf0, 0xeb, 0x70, 0xf1, 0x8c, 0xd7, 0x27, 0x72, 0xec, 0x0d, 0x08, 0xce, 0x5f, 0x10,
	0xd2, 0x7c, 0xee, 0xf8, 0xcd, 0xe2, 0x68, 0xb5, 0x76, 0xef, 0x39, 0xd0, 0x10, 0x11, 0x4f, 0x84,
	0xc4, 0x7f, 0x83, 0x8f, 0xb2, 0xac, 0x36, 0xed, 0x0c, 0xff, 0xdd, 0x21, 0x27, 0xa1, 0xe3, 0x4d,
	0x6f, 0xca, 0x4b, 0xaa, 0x5b, 0x54, 0xb3, 0xc7, 0x22, 0x63, 0xfc, 0x1f, 0x10, 0xf5, 0x0a, 0x59,
	0x4d, 0xf2, 0x79, 0xd8, 0xc4, 0x3f, 0x21, 0x87, 0x8d, 0x34, 0x4c, 0x78, 0x2f, 0x64, 0x25, 0x3d,
	0x14, 0xf3, 0x5f, 0xe0, 0x74, 0x95, 0x03, 0x31, 0xdb, 0xe7, 0x01, 0xc3, 0xff, 0x06, 0xa7, 0x59,
	0x15, 0x5b, 0x92, 0xd5, 0x24, 0x5f, 0x62, 0x72, 0xf8, 0x64, 0xfe, 0x07, 0x3e, 0xdb, 0x41, 0x1d,
	0xf8, 0xfb, 0x2f, 0x34, 0xaf, 0xfa, 0x1e, 0xe3, 0xf1, 0x80, 0x08, 0x13, 0xda, 0x29, 0xf6, 0xf4,
	0x56, 0x49, 0x63, 0x0e, 0xe1, 0xb0, 0xe6, 0xed, 0x92, 0xa6, 0x2e, 0xa2, 0x8e, 0xba, 0xb7, 0x86,
	0x35, 0xef, 0x40, 0x19, 0xea, 0x7b, 0x0c, 0xd7, 0x67, 0xf8, 0xfc, 0x8a, 0x49, 0x4c, 0x19, 0x9e,
	0x61, 0xed, 0x34, 0x60, 0x93, 0x92, 0x1d, 0x4e, 0xb3, 0x39, 0xec, 0x82, 0x8a, 0x29, 0xd6, 0x4a,
	0x07, 0x7d, 0x86, 0x3f, 0x55, 0x59, 0x19, 0xc7, 0xf0, 0x87, 0xf2, 0xe9, 0x0a, 0xf9, 0x00, 0xda,
	0xf5, 0xae, 0x1a, 0x9d, 0xdd, 0xcf, 0x54, 0x56, 0x06, 0x3c, 0x6c, 0xec, 0xb3, 0x85, 0x66, 0x6c,
	0xb5, 0xc4, 0x7d, 0xae, 0xa4, 0x59, 0x2d, 0x71, 0x9f, 0x2f, 0x69, 0x56, 0x4b, 0xdc, 0x17, 0x2a,
	0x90, 0xb8, 0xb1, 0x52, 0xcc, 0x5f, 0x2c, 0x12, 0x37, 0xb6, 0x7a, 0xe2, 0xbe, 0x54, 0x24, 0x6e,
	0xcc, 0x92, 0xb8, 0x2f, 0x5b, 0xe2, 0x18, 0x8e, 0xf5, 0x2b, 0x45, 0xe2, 0xc6, 0xde, 0x35, 0x71,
	0x5f, 0xb5, 0x04, 0x3c, 0x6c, 0xec, 0x6b, 0x15, 0xd5, 0x96, 0xc6, 0x1b, 0xfe, 0x21, 0x8f, 0x85,
	0xa1, 0x19, 0x4c, 0xbe, 0x9e, 0x85, 0x58, 0x46, 0xb3, 0xb1, 0xeb, 0x1b, 0x15, 0x35, 0x01, 0x64,
	0xf0, 0xca, 0x76, 0xfe, 0xcd, 0x8a, 0xba, 0x1a, 0x56, 0x90, 0x30, 0x36, 0x7d, 0xab, 0xa2, 0x1a,
	0xb6, 0xe2, 0xc7, 0x69, 0x92, 0x30, 0x35, 0xd9, 0x2f, 0x31, 0x19, 0xb3, 0x36, 0xfe, 0x76, 0x45,
	0xb5, 0xcc, 0x12, 0xd3, 0xe0, 0xb1, 0xc2, 0xbf, 0x53, 0x51, 0x8d, 0xbc, 0x84, 0x9b, 0xed, 0x7d,
	0xb7, 0xa2, 0xee, 0x08, 0x45, 0x28, 0xc4, 0x6f, 0x7a, 0xf8, 0x7b, 0x60, 0x7c, 0x78, 0x36, 0xfd,
	0x3e, 0x18, 0x19, 0x1a, 0x72, 0x7f, 0x50, 0x51, 0x4d, 0xce, 0x18, 0x99, 0x38, 0x92, 0x30, 0x19,
	0x51, 0x73, 0x20, 0xf0, 0x0f, 0x87, 0x77, 0x0b, 0xe3, 0xdc, 0x8f, 0x2a, 0xea, 0x0a, 0x28, 0x31,
	0xf9, 0x92, 0x62, 0x44, 0xf9, 0x71, 0x56, 0x63, 0x93, 0xbd, 0x9c, 0xae, 0xf1, 0xf9, 0x79, 0xed,
	0xf8, 0x27, 0x15, 0xd5, 0x70, 0x33, 0x03, 0xe6, 0x3e, 0xff, 0x69, 0x56, 0x29, 0xb3, 0xc2, 0xa7,
	0xdd, 0x5e, 0xa8, 0x86, 0x3b, 0xc1, 0xa3, 0xe2, 0x39, 0xf5, 0xb3, 0x0a, 0x59, 0x8b, 0xdc, 0x96,
	0xaf, 0x06, 0xe9, 0x6c, 0xd4, 0x77, 0x55, 0x8d, 0xd4, 0xff, 0x55, 0x2f, 0x9b, 0x53, 0xf2, 0x69,
	0x1f, 0xe0, 0x81, 0x49, 0xf7, 0x42, 0x57, 0xe5, 0x18, 0x60, 0x88, 0xe6, 0x22, 0x57, 0x4d, 0xd9,
	0x80, 0xd7, 0x69, 0xb0, 0xa8, 0x3c, 0x2f, 0xd0, 0x98, 0xe1, 0x8b, 0x5d, 0x75, 0x36, 0x14, 0xa7,
	0x47, 0xa9, 0xac, 0x2f, 0xe3, 0x4b, 0xc0, 0xd2, 0xe0, 0xc8, 0x7a, 0xa9, 0x05, 0x57, 0x1e, 0x2e,
	0x73, 0x55, 0x01, 0x0a, 0x1c, 0x5c, 0x5f, 0x3e, 0x64, 0x5e, 0xef, 0xff, 0x0a, 0x57, 0x65, 0x27,
	0x43, 0x61, 0xda, 0x77, 0xd5, 0xb1, 0x53, 0x90, 0xe5, 0xd8, 0x5d, 0x09, 0x7a, 0x78, 0x5f, 0x5c,
	0x05, 0x1e, 0xbd, 0x50, 0xc4, 0xa5, 0xc2, 0x5c, 0xed, 0xaa, 0x2f, 0x47, 0x11, 0x33, 0x2c, 0x48,
	0xf8, 0x3c, 0x67, 0x72, 0x7f, 0xf3, 0xc3, 0x16, 0xa3, 0xd7, 0xd8, 0x84, 0xbb, 0x2d, 0xc2, 0x6b,
	0x6d, 0xc2, 0x3d, 0x16, 0xe1, 0x75, 0x36, 0xa1, 0x6d, 0x72, 0xbb, 0xde, 0x26, 0xdc, 0x6b, 0x11,
	0xde, 0x60, 0x13, 0x9e, 0x6a, 0x11, 0xde, 0x68, 0x13, 0x9e, 0x66, 0x11, 0xde, 0x64, 0x13, 0x9e,
	0x6e, 0x11, 0xde, 0x6c, 0x13, 0x9e, 0x61, 0x11, 0xde, 0x62, 0x13, 0x9e, 0x69, 0x7b, 0x8b, 0xb8,
	0xe4, 0x83, 0xe8, 0xc4, 0xe1, 0x84, 0xdb, 0x4a, 0x73, 0xab, 0x55, 0x69, 0xab, 0xcd, 0x6d, 0x56,
	0xa5, 0xad, 0x38, 0xb7, 0xc3, 0x47, 0x50, 0x3e, 0x17, 0x9a, 0xbb, 0xc3, 0xc6, 0xe5, 0xaf, 0xd6,
	0x3b, 0x6d, 0xdc, 0x1e, 0xcd, 0xdd, 0x65, 0xe3, 0xc6, 0x34, 0x77, 0xb7, 0x8d, 0xdb, 0xab, 0xb9,
	0x7b, 0x6c, 0xdc, 0xa9, 0x9a, 0xbb, 0xd7, 0xc6, 0x9d, 0xa6, 0xb9, 0xfb, 0x6c, 0xdc, 0xe9, 0x9a,
	0xbb, 0xdf, 0xc6, 0x9d, 0xa1, 0xb9, 0x07, 0x6c, 0xdc, 0x99, 0x9a, 0xfb, 0x03, 0x7c, 0x78, 0x03,
	0x55, 0xc9, 0x9f, 0x26, 0x56, 0x32, 0xcf, 0xcc, 0x83, 0x56, 0x32, 0x4f, 0xcd, 0x43, 0xae, 0x9a,
	0xaa, 0x74, 0xcf, 0x31, 0xcd, 0xda, 0x52, 0x95, 0x87, 0x5d, 0x35, 0x55, 0x29, 0x59, 0xf6, 0x96,
	0xf0, 0x03, 0xc9, 0x98, 0x4d, 0xf7, 0x08, 0x74, 0xbc, 0x4c, 0x37, 0xdd, 0xcc, 0x9f, 0x2a, 0xae,
	0xea, 0xc2, 0xa5, 0xfe, 0xa8, 0xfa, 0x49, 0x71, 0x11, 0x3f, 0xea, 0xaa, 0xeb, 0xaa, 0x44, 0x67,
	0x7d, 0xa9, 0xe0, 0x1f, 0x83, 0x76, 0x36, 0xbb, 0xdc, 0x63, 0x35, 0x1e, 0xd3, 0x20, 0x10, 0xb2,
	0x8d, 0x1f, 0x77, 0xd5, 0x35, 0xaf, 0xf0, 0x29, 0x29, 0xd2, 0x9e, 0x99, 0x4f, 0xf5, 0xc8, 0x0b,
	0x2d, 0xeb, 0x09, 0x97, 0x8c, 0xa0, 0x0f, 0xad, 0xa2, 0xb2, 0xc4, 0xf0, 0x67, 0x68, 0x59, 0xd3,
	0x11, 0x4f, 0x38, 0x0d, 0xf9, 0x79, 0xfa, 0x81, 0xf8, 0x64, 0xb9, 0x49, 0x66, 0x3f, 0x98, 0x29,
	0xf9, 0x53, 0xd0, 0xf4, 0xe0, 0x12, 0x7d, 0xda, 0x55, 0xd7, 0x19, 0x34, 0xf2, 0xac, 0x83, 0x4f,
	0x77, 0xcf, 0xa1, 0x21, 0x8d, 0x02, 0x86, 0x9f, 0x81, 0x2a, 0xb7, 0xa4, 0x48, 0x54, 0x49, 0x8a,
	0x5f, 0x68, 0x9e, 0x75, 0xd5, 0x55, 0x37, 0xc8, 0x69, 0x37, 0xcf, 0xb9, 0x6a, 0xd2, 0x87, 0xcc,
	0xd6, 0x79, 0x97, 0x27, 0xf8, 0x79, 0x00, 0xe7, 0x22, 0xb6, 0xc4, 0xa2, 0xc9, 0x50, 0xf4, 0xf1,
	0x0b, 0xe0, 0x3d, 0x1f, 0x30, 0x32, 0xad, 0xf1, 0xf0, 0x22, 0x78, 0x2f, 0xb3, 0xda, 0xc7, 0x4b,
	0xae, 0xba, 0x4e, 0xb3, 0x00, 0x57, 0x99, 0xdb, 0x5e, 0x86, 0xb3, 0x64, 0x14, 0xcd, 0xbe, 0x99,
	0x99, 0xf0, 0x2b, 0x40, 0x9a, 0x09, 0xa7, 0x44, 0xbe, 0x0a, 0xb6, 0x57, 0x9b, 0xe5, 0x5e, 0x83,
	0xe5, 0x96, 0x61, 0xec, 0x75, 0xd8, 0xf6, 0x44, 0xc8, 0x82, 0x44, 0xf2, 0xa0, 0x38, 0x1e, 0x6f,
	0x80, 0xe9, 0xd5, 0x46, 0xc9, 0x37, 0x5d, 0x75, 0xa7, 0x2b, 0xc5, 0xea, 0xd3, 0xe0, 0x5f, 0x21,
	0x97, 0x25, 0xb7, 0x6f, 0x41, 0x2e, 0x5b, 0x8d, 0x69, 0x2f, 0x7b, 0x7d, 0xb5, 0x8b, 0xd7, 0xd7,
	0xdb, 0xae, 0x7a, 0xd9, 0x2b, 0xd6, 0x3c, 0x5c, 0x56, 0xcc, 0x1f, 0xef, 0xc0, 0x99, 0x99, 0xf1,
	0x6a, 0x87, 0xaa, 0xde, 0xa1, 0xfa, 0x44, 0xf5, 0x2c, 0x7c, 0x3e, 0x2a, 0xa3, 0xb5, 0x1c, 0xbd,
	0x00, 0x8d, 0xef, 0xdb, 0x57, 0xf9, 0x48, 0xad, 0xc3, 0x93, 0x20, 0xee, 0xc6, 0x23, 0xe7, 0x8a,
	0xe5, 0x58, 0x44, 0x87, 0x53, 0x1e, 0x8d, 0x04, 0xa2, 0x3b, 0xaa, 0xb0, 0xd1, 0x5e, 0xfe, 0x1b,
	0xfb, 0x68, 0x47, 0x84, 0x34, 0xea, 0x8c, 0x2e, 0x50, 0xd9, 0xee, 0x53, 0xc9, 0x46, 0x87, 0x7e,
	0x99, 0x3f, 0xe7, 0x3d, 0x99, 0x72, 0xec, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x6d, 0x4d, 0xd1,
	0x7d, 0xb3, 0x17, 0x00, 0x00,
}
