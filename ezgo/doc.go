// Package ezgo 安全携程模块
/*
	示例:
	safeGo := ezgo.New[string](func(args ...any) {
		fmt.Println(args)
	})
	safeGo.SetGoBeforeHandler(func() string {
		return ezgin.MWTrace.GetCurRoutineId()
	})
	safeGo.SetGoAfterHandler(func(params string) {
		ezgin.MWTrace.CopyPreHeaderToCurRoutine(params)
	})
	safeGo.Run("hello", "world")
*/
package ezgo
