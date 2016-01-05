#ifdef __cplusplus
extern "C" {
#endif

void* QScriptClass_NewQScriptClass(void* engine);
void* QScriptClass_Engine(void* ptr);
void* QScriptClass_Extension(void* ptr, int extension, void* argument);
void* QScriptClass_ExtensionDefault(void* ptr, int extension, void* argument);
char* QScriptClass_Name(void* ptr);
char* QScriptClass_NameDefault(void* ptr);
void* QScriptClass_NewIterator(void* ptr, void* object);
void* QScriptClass_NewIteratorDefault(void* ptr, void* object);
void* QScriptClass_Prototype(void* ptr);
void* QScriptClass_PrototypeDefault(void* ptr);
int QScriptClass_SupportsExtension(void* ptr, int extension);
int QScriptClass_SupportsExtensionDefault(void* ptr, int extension);
void QScriptClass_DestroyQScriptClass(void* ptr);
char* QScriptClass_ObjectNameAbs(void* ptr);
void QScriptClass_SetObjectNameAbs(void* ptr, char* name);
void* QScriptContext_ActivationObject(void* ptr);
void* QScriptContext_Argument(void* ptr, int index);
int QScriptContext_ArgumentCount(void* ptr);
void* QScriptContext_ArgumentsObject(void* ptr);
char* QScriptContext_Backtrace(void* ptr);
void* QScriptContext_Callee(void* ptr);
void* QScriptContext_Engine(void* ptr);
int QScriptContext_IsCalledAsConstructor(void* ptr);
void* QScriptContext_ParentContext(void* ptr);
void QScriptContext_SetActivationObject(void* ptr, void* activation);
void QScriptContext_SetThisObject(void* ptr, void* thisObject);
int QScriptContext_State(void* ptr);
void* QScriptContext_ThisObject(void* ptr);
void* QScriptContext_ThrowError(void* ptr, int error, char* text);
void* QScriptContext_ThrowError2(void* ptr, char* text);
void* QScriptContext_ThrowValue(void* ptr, void* value);
char* QScriptContext_ToString(void* ptr);
void QScriptContext_DestroyQScriptContext(void* ptr);
void* QScriptContextInfo_NewQScriptContextInfo3();
void* QScriptContextInfo_NewQScriptContextInfo(void* context);
void* QScriptContextInfo_NewQScriptContextInfo2(void* other);
char* QScriptContextInfo_FileName(void* ptr);
int QScriptContextInfo_FunctionEndLineNumber(void* ptr);
int QScriptContextInfo_FunctionMetaIndex(void* ptr);
char* QScriptContextInfo_FunctionName(void* ptr);
char* QScriptContextInfo_FunctionParameterNames(void* ptr);
int QScriptContextInfo_FunctionStartLineNumber(void* ptr);
int QScriptContextInfo_FunctionType(void* ptr);
int QScriptContextInfo_IsNull(void* ptr);
int QScriptContextInfo_LineNumber(void* ptr);
long long QScriptContextInfo_ScriptId(void* ptr);
void QScriptContextInfo_DestroyQScriptContextInfo(void* ptr);
void* QScriptEngine_NewQScriptEngine();
void* QScriptEngine_NewQScriptEngine2(void* parent);
void QScriptEngine_AbortEvaluation(void* ptr, void* result);
void* QScriptEngine_Agent(void* ptr);
char* QScriptEngine_AvailableExtensions(void* ptr);
void QScriptEngine_ClearExceptions(void* ptr);
void QScriptEngine_CollectGarbage(void* ptr);
void* QScriptEngine_CurrentContext(void* ptr);
void* QScriptEngine_DefaultPrototype(void* ptr, int metaTypeId);
void* QScriptEngine_Evaluate2(void* ptr, void* program);
void* QScriptEngine_Evaluate(void* ptr, char* program, char* fileName, int lineNumber);
void* QScriptEngine_GlobalObject(void* ptr);
int QScriptEngine_HasUncaughtException(void* ptr);
void* QScriptEngine_ImportExtension(void* ptr, char* extension);
char* QScriptEngine_ImportedExtensions(void* ptr);
void QScriptEngine_InstallTranslatorFunctions(void* ptr, void* object);
int QScriptEngine_IsEvaluating(void* ptr);
void* QScriptEngine_NewDate2(void* ptr, void* value);
void* QScriptEngine_NewObject(void* ptr);
void* QScriptEngine_NewObject2(void* ptr, void* scriptClass, void* data);
void* QScriptEngine_NewQMetaObject(void* ptr, void* metaObject, void* ctor);
void* QScriptEngine_NewQObject(void* ptr, void* object, int ownership, int options);
void* QScriptEngine_NewQObject2(void* ptr, void* scriptObject, void* qtObject, int ownership, int options);
void* QScriptEngine_NewRegExp(void* ptr, void* regexp);
void* QScriptEngine_NewRegExp2(void* ptr, char* pattern, char* flags);
void* QScriptEngine_NewVariant2(void* ptr, void* object, void* value);
void* QScriptEngine_NewVariant(void* ptr, void* value);
void* QScriptEngine_NullValue(void* ptr);
void QScriptEngine_PopContext(void* ptr);
int QScriptEngine_ProcessEventsInterval(void* ptr);
void* QScriptEngine_PushContext(void* ptr);
void QScriptEngine_ReportAdditionalMemoryCost(void* ptr, int size);
void QScriptEngine_SetAgent(void* ptr, void* agent);
void QScriptEngine_SetDefaultPrototype(void* ptr, int metaTypeId, void* prototype);
void QScriptEngine_SetGlobalObject(void* ptr, void* object);
void QScriptEngine_SetProcessEventsInterval(void* ptr, int interval);
void QScriptEngine_ConnectSignalHandlerException(void* ptr);
void QScriptEngine_DisconnectSignalHandlerException(void* ptr);
void QScriptEngine_SignalHandlerException(void* ptr, void* exception);
void* QScriptEngine_ToObject(void* ptr, void* value);
void* QScriptEngine_UncaughtException(void* ptr);
char* QScriptEngine_UncaughtExceptionBacktrace(void* ptr);
int QScriptEngine_UncaughtExceptionLineNumber(void* ptr);
void* QScriptEngine_UndefinedValue(void* ptr);
void QScriptEngine_DestroyQScriptEngine(void* ptr);
void QScriptEngine_TimerEvent(void* ptr, void* event);
void QScriptEngine_TimerEventDefault(void* ptr, void* event);
void QScriptEngine_ChildEvent(void* ptr, void* event);
void QScriptEngine_ChildEventDefault(void* ptr, void* event);
void QScriptEngine_CustomEvent(void* ptr, void* event);
void QScriptEngine_CustomEventDefault(void* ptr, void* event);
void* QScriptEngineAgent_NewQScriptEngineAgent(void* engine);
void QScriptEngineAgent_ContextPop(void* ptr);
void QScriptEngineAgent_ContextPopDefault(void* ptr);
void QScriptEngineAgent_ContextPush(void* ptr);
void QScriptEngineAgent_ContextPushDefault(void* ptr);
void* QScriptEngineAgent_Engine(void* ptr);
void QScriptEngineAgent_ExceptionCatch(void* ptr, long long scriptId, void* exception);
void QScriptEngineAgent_ExceptionCatchDefault(void* ptr, long long scriptId, void* exception);
void QScriptEngineAgent_ExceptionThrow(void* ptr, long long scriptId, void* exception, int hasHandler);
void QScriptEngineAgent_ExceptionThrowDefault(void* ptr, long long scriptId, void* exception, int hasHandler);
void* QScriptEngineAgent_Extension(void* ptr, int extension, void* argument);
void* QScriptEngineAgent_ExtensionDefault(void* ptr, int extension, void* argument);
void QScriptEngineAgent_FunctionEntry(void* ptr, long long scriptId);
void QScriptEngineAgent_FunctionEntryDefault(void* ptr, long long scriptId);
void QScriptEngineAgent_FunctionExit(void* ptr, long long scriptId, void* returnValue);
void QScriptEngineAgent_FunctionExitDefault(void* ptr, long long scriptId, void* returnValue);
void QScriptEngineAgent_PositionChange(void* ptr, long long scriptId, int lineNumber, int columnNumber);
void QScriptEngineAgent_PositionChangeDefault(void* ptr, long long scriptId, int lineNumber, int columnNumber);
void QScriptEngineAgent_ScriptLoad(void* ptr, long long id, char* program, char* fileName, int baseLineNumber);
void QScriptEngineAgent_ScriptLoadDefault(void* ptr, long long id, char* program, char* fileName, int baseLineNumber);
void QScriptEngineAgent_ScriptUnload(void* ptr, long long id);
void QScriptEngineAgent_ScriptUnloadDefault(void* ptr, long long id);
int QScriptEngineAgent_SupportsExtension(void* ptr, int extension);
int QScriptEngineAgent_SupportsExtensionDefault(void* ptr, int extension);
void QScriptEngineAgent_DestroyQScriptEngineAgent(void* ptr);
char* QScriptEngineAgent_ObjectNameAbs(void* ptr);
void QScriptEngineAgent_SetObjectNameAbs(void* ptr, char* name);
void QScriptExtensionPlugin_Initialize(void* ptr, char* key, void* engine);
char* QScriptExtensionPlugin_Keys(void* ptr);
void* QScriptExtensionPlugin_SetupPackage(void* ptr, char* key, void* engine);
void QScriptExtensionPlugin_DestroyQScriptExtensionPlugin(void* ptr);
void QScriptExtensionPlugin_TimerEvent(void* ptr, void* event);
void QScriptExtensionPlugin_TimerEventDefault(void* ptr, void* event);
void QScriptExtensionPlugin_ChildEvent(void* ptr, void* event);
void QScriptExtensionPlugin_ChildEventDefault(void* ptr, void* event);
void QScriptExtensionPlugin_CustomEvent(void* ptr, void* event);
void QScriptExtensionPlugin_CustomEventDefault(void* ptr, void* event);
void* QScriptProgram_NewQScriptProgram();
void* QScriptProgram_NewQScriptProgram3(void* other);
void* QScriptProgram_NewQScriptProgram2(char* sourceCode, char* fileName, int firstLineNumber);
char* QScriptProgram_FileName(void* ptr);
int QScriptProgram_FirstLineNumber(void* ptr);
int QScriptProgram_IsNull(void* ptr);
char* QScriptProgram_SourceCode(void* ptr);
void QScriptProgram_DestroyQScriptProgram(void* ptr);
void* QScriptString_NewQScriptString();
void* QScriptString_NewQScriptString2(void* other);
int QScriptString_IsValid(void* ptr);
char* QScriptString_ToString(void* ptr);
void QScriptString_DestroyQScriptString(void* ptr);
void* QScriptSyntaxCheckResult_NewQScriptSyntaxCheckResult(void* other);
int QScriptSyntaxCheckResult_ErrorColumnNumber(void* ptr);
int QScriptSyntaxCheckResult_ErrorLineNumber(void* ptr);
char* QScriptSyntaxCheckResult_ErrorMessage(void* ptr);
int QScriptSyntaxCheckResult_State(void* ptr);
void QScriptSyntaxCheckResult_DestroyQScriptSyntaxCheckResult(void* ptr);
void* QScriptValue_NewQScriptValue();
void* QScriptValue_NewQScriptValue10(int value);
void* QScriptValue_NewQScriptValue11(int value);
void* QScriptValue_NewQScriptValue16(void* value);
void* QScriptValue_NewQScriptValue2(void* other);
void* QScriptValue_NewQScriptValue15(char* value);
void* QScriptValue_NewQScriptValue17(char* value);
void* QScriptValue_NewQScriptValue12(int value);
void* QScriptValue_Call2(void* ptr, void* thisObject, void* arguments);
void* QScriptValue_Construct2(void* ptr, void* arguments);
void* QScriptValue_Data(void* ptr);
void* QScriptValue_Engine(void* ptr);
int QScriptValue_Equals(void* ptr, void* other);
int QScriptValue_InstanceOf(void* ptr, void* other);
int QScriptValue_IsArray(void* ptr);
int QScriptValue_IsBool(void* ptr);
int QScriptValue_IsDate(void* ptr);
int QScriptValue_IsError(void* ptr);
int QScriptValue_IsFunction(void* ptr);
int QScriptValue_IsNull(void* ptr);
int QScriptValue_IsNumber(void* ptr);
int QScriptValue_IsObject(void* ptr);
int QScriptValue_IsQMetaObject(void* ptr);
int QScriptValue_IsQObject(void* ptr);
int QScriptValue_IsRegExp(void* ptr);
int QScriptValue_IsString(void* ptr);
int QScriptValue_IsUndefined(void* ptr);
int QScriptValue_IsValid(void* ptr);
int QScriptValue_IsVariant(void* ptr);
int QScriptValue_LessThan(void* ptr, void* other);
void* QScriptValue_Property2(void* ptr, void* name, int mode);
void* QScriptValue_Property(void* ptr, char* name, int mode);
int QScriptValue_PropertyFlags2(void* ptr, void* name, int mode);
int QScriptValue_PropertyFlags(void* ptr, char* name, int mode);
void* QScriptValue_Prototype(void* ptr);
void* QScriptValue_ScriptClass(void* ptr);
void QScriptValue_SetData(void* ptr, void* data);
void QScriptValue_SetProperty2(void* ptr, void* name, void* value, int flags);
void QScriptValue_SetProperty(void* ptr, char* name, void* value, int flags);
void QScriptValue_SetPrototype(void* ptr, void* prototype);
void QScriptValue_SetScriptClass(void* ptr, void* scriptClass);
int QScriptValue_StrictlyEquals(void* ptr, void* other);
int QScriptValue_ToBool(void* ptr);
void* QScriptValue_ToDateTime(void* ptr);
void* QScriptValue_ToQMetaObject(void* ptr);
void* QScriptValue_ToQObject(void* ptr);
void* QScriptValue_ToRegExp(void* ptr);
char* QScriptValue_ToString(void* ptr);
void* QScriptValue_ToVariant(void* ptr);
void QScriptValue_DestroyQScriptValue(void* ptr);
void* QScriptable_Argument(void* ptr, int index);
int QScriptable_ArgumentCount(void* ptr);
void* QScriptable_Context(void* ptr);
void* QScriptable_Engine(void* ptr);
void* QScriptable_ThisObject(void* ptr);

#ifdef __cplusplus
}
#endif