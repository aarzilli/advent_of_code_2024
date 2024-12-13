#!/usr/bin/env wolframscript
Off[Minimize::infeas]
Print@Total[#[[1]]& /@
		Select[ToExpression /@ Import["13.example.preproc", "Lines"], 
			#[[1]] != Infinity &]]
Print@Total[#[[1]]& /@ 
		Select[ToExpression /@ Import["13.txt.preproc", "Lines"], 
			#[[1]] != Infinity &]]
Print@Total[#[[1]]& /@
		Select[ToExpression /@ Import["13.txt.preproc2", "Lines"],
			#[[1]] != Infinity&]]



