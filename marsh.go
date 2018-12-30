package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Ans struct {
	Count   int `json:"count"`
	Results []struct {
		FileName   string `json:"fileName"`
		Path       string `json:"path"`
		Collection struct {
			Name string `json:"name"`
		} `json:"collection"`
		Project struct {
			ID         string `json:"id"`
			Name       string `json:"name"`
			Visibility string `json:"visibility"`
		} `json:"project"`
		Wiki struct {
			Name       string `json:"name"`
			ID         string `json:"id"`
			MappedPath string `json:"mappedPath"`
			Version    string `json:"version"`
		} `json:"wiki"`
		ContentID string `json:"contentId"`
		Hits      []struct {
			FieldReferenceName string   `json:"fieldReferenceName"`
			Highlights         []string `json:"highlights"`
		} `json:"hits"`
	} `json:"results"`
	InfoCode int `json:"infoCode"`
	Facets   struct {
		Project []struct {
			Name        string `json:"name"`
			ID          string `json:"id"`
			ResultCount int    `json:"resultCount"`
		} `json:"Project"`
		Wiki []struct {
			Name        string `json:"name"`
			ID          string `json:"id"`
			ResultCount int    `json:"resultCount"`
		} `json:"Wiki"`
	} `json:"facets"`
}

type Message struct {
	Type string `json:"type"`
	From struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"from"`
	Conversation struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"conversation"`
	Recipient struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"recipient"`
	Text      string `json:"text"`
	ReplyToID string `json:"replyToId"`
}

type Actioncard struct {
	Context         string `json:"@context"`
	Type            string `json:"@type"`
	ThemeColor      string `json:"themeColor"`
	Title           string `json:"title"`
	Text            string `json:"text"`
	PotentialAction []struct {
		Type   string `json:"@type"`
		Name   string `json:"name"`
		Inputs []struct {
			Type        string `json:"@type"`
			ID          string `json:"id"`
			IsMultiline bool   `json:"isMultiline"`
			Title       string `json:"title"`
		} `json:"inputs,omitempty"`
		Actions []struct {
			Type      string `json:"@type"`
			Name      string `json:"name"`
			IsPrimary bool   `json:"isPrimary"`
			Target    string `json:"target"`
		} `json:"actions,omitempty"`
		Targets []struct {
			Os  string `json:"os"`
			URI string `json:"uri"`
		} `json:"targets,omitempty"`
	} `json:"potentialAction"`
}

func Rt(body []byte) (*Ans, error) {

	var s = new(Ans)
	err := json.Unmarshal(body, &s)
	if err != nil {
		fmt.Println(err)
	}
	return s, err

}

func Replacer(j []byte, t string) (re []byte) {
	s := string(j)
	result := strings.Replace(s, "kubernetes", t, -1)
	return []byte(result)

}

func Collector(s *Ans) (k []string) {
	g := make([]string, 0)
	for _, re := range s.Results {
		fmt.Println("RESULT ************ ", re.Path)
		for _, tf := range re.Hits {

			for _, k := range tf.Highlights {
				result := strings.Replace(k, "<highlighthit>", "", -1)
				result1 := strings.Replace(result, "</highlighthit>", "", -1)
				//fmt.Println(result1)
				separator := "\n"
				g = append(g, separator, result1, separator, re.Path, separator)

			}

		}

	}
	return g
}

var jsonStr = []byte(`{"searchText": "kubernetes","filters": {
	"Project": [
	  "Wiki"
	  
	]
  },"$top": 6,"includeFacets": true,"WikiResult": ,}`)

func Devopsresp(m Message, st string) (jo string) {
	ret := Replacer(jsonStr, st)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(ret))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Basic *************************************")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println("response Body:", string(body))
	s, _ := Rt(body)
	resu := Collector(s)
	var joined = strings.Join(resu, "")
	//fmt.Println(resu)
	return joined

}
