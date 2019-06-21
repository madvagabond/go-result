package result


import (
	"testing"
	"errors"
)



const (
	hello  = "hey you"
	expect = "hey you out there in the cold"
)


var (
	test_err = 	errors.New("test error")
)



func TestResult(t *testing.T) {

	
	ok :=  Ok(hello)


	
	map_res := ok.Map(func(interface{}) interface{} {
		return expect
	})


	out, e := map_res.Unwrap()


	if e != nil {
		t.Error(e) 
	}
	




	if out.(string) != expect {
		t.Errorf("unexpected value %s", out.(string))
	}



	bind_res := map_res.Bind(func(x interface{}) Result {

		return Wrap(x, nil)
	})


	
	if bind_res.IsOk() == false {
		t.Error("result unsuccessful")
	}



	err_result := Err(test_err).Map(func(interface{}) interface{} {
		return expect
	})
	

	val, err := err_result.Unwrap()


	if err_result.IsErr() == false {
		err := errors.New("result doesn't return error")
		t.Error(err)
	}


	
	if val != nil {
		t.Error("Map function was performed on error")
	}


	if err == nil {
		t.Errorf("Error is empty in failed result")
	}

	
}



