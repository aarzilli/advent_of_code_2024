#!/usr/bin/env wolframscript
Off[Minimize::infeas]
Print@Total[#[[1]]& /@ 
		Select[ToExpression /@ Import[$ScriptCommandLine[[2]], "Lines"], 
			#[[1]] != Infinity &]]



