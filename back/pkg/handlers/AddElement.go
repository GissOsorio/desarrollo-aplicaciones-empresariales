package handlers

// import (

// )

// func AddElement(w http.ResponseWriter, r *http.Request) {
// 	// Read to request body
// 	defer r.Body.Close()
// 	body, err := ioutil.ReadAll(r.Body)

// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	var element models.Element
// 	json.Unmarshal(body, &element)

// 	element.Id = (uuid.New()).String()

// 	mocks.Elements = append(mocks.Elements, element)

// 	w.Header().Add("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusCreated)
// 	json.NewEncoder(w).Encode("Created")
// }