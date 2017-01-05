package utils

func Trycatch(try func(), catch func(e interface{})) {
	defer func() {
		if err := recover(); err != nil {
			catch(err)
		}
	}()
	try()
}
