package result



type Result struct {
	value interface{}
	err error
}



func Ok(val interface{}) Result {
	return Result {value: val, err: nil}
}



func Err(e error) Result {
	return Result {value: nil, err: e}
}



func Wrap(value interface{}, err error) Result {
	return Result {value, err }
}


func (R Result) Map( fun func(interface{}) interface{} ) Result {

	if R.err != nil {
		return R 
	}

	return Result {value: fun(R.value), err: R.err }
}





func (R Result) Bind(fun func(interface{}) Result ) Result {

	if R.err != nil {
		return R
	}


	return fun(R.value)

}



func (R Result) IsOk() bool {
	return R.err == nil
}



func (R Result) IsErr() bool {
	return R.err != nil 
}





func (R Result) Unwrap() (interface{}, error) {
	return R.value, R.err
}
