package handler


func HandlerSuccess(res http.ResponseWriter, req *http.Request) {

	user, err := gothic.CompleteUserAuth(res, req)
	if err != nil {
		fmt.Fprintln(res, err)
		return
	}
	t, _ := template.ParseFiles("templates/success.html")
	t.Execute(res, user)
}