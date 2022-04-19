function()
{
    HRESULT error = S_OK

    if !(SUCCEEDED(Operation1())){
		error = OPERATION1FAILED
		return error
	}
	
	if !(SUCCEEDED(Operation2())){
		error = OPERATION2FAILED
		return error
	}
	
	if !(SUCCEEDED(Operation3())){
		error = OPERATION3FAILED
		return error
	}
	
	if !(SUCCEEDED(Operation4())){
		error = OPERATION4FAILED
		return error
	}

    return error
}
