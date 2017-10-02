package crawler


func TestCrawl(t *testing.T){
	type fakeFetcher map[string]*fakeResult
	type fakeResult struct {
		body string
		urls []string
	}

	func (f fakeFetcher) Fetch(url string) (string, []string, error) {
		if val, ok := f[url]; ok {
			return val.body, val.urls, nil
		}
		return "", nil, fmt.Errorf("not found: %s", url)
	}
	var fetcher = fakeFetcher{
		"http://golang.org/": &fakeResult{
			"Go Website",
			[]string{
				"http://golang.org/pkg/",
				"http://golang.org/os/",
			},
		},
		"http://golang.org/pkg/": &fakeResult{
			"Packages",
			[]string{
				"http://golang.org/pkg/os/",
				"http://golang.org/pkg/fmt/",
				"http://golang.org/",
			},
		},
		"http://golang.org/os/": &fakeResult{
			"OS stuff",
			[]string{
				"http://golang.org/os/detect",
				"http://golang.org/pkg",
				"http://golang.org/",
			},
		},
	}

	go func(){
		
	}()
}
