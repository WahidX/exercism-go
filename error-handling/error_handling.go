package erratum


func Use(o ResourceOpener, input string) (err error) {
	resource, err := o()

	if err != nil {
		if _, ok := err.(TransientError); ok {
			return Use(o, input)
		}
		return
	}

	defer resource.Close()

	defer func() {
		if r := recover(); r!=nil {
			switch e := r.(type) {
			case FrobError:
				resource.Defrob(e.defrobTag)
				err = e
			case error:
				err = e	
			}
		}
	}()

	resource.Frob(input)

	return
}




