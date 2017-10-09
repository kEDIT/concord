package crawler


//Rendition 2

//Mock Fetcher
/*
 *type fakeFetcher struct{}
 *type mockUrl struct{
 *    parent string
 *    children []string
 *    depth int
 *}
 *
 *
 *func TestCrawl(t *testing.T) {
 *    
 *    var cases = []struct{
 *        m mockUrl
 *    }{
 *         {
 *             m: &mockUrl{
 *                parent: "http://testurl.com/",
 *                children: []string{
 *                    "http://uberurl.com/",
 *                    "https://othertesturl.org",
 *                    "fithub.com/",
 *                    "ecommerce",
 *                },
 *                depth: 0,
 *            },
 *        },
 *        {
 *            m: &mockUrl{
 *                parent: "http://uberurl.com/",
 *                children: []string{
 *                    "fithub.com/fits",
 *                    "email@email.org",
 *                    "image-link-to-ecommerce",
 *                },
 *                depth: 1,
 *            },
 *        },
 *        {
 *            m: &mockUrl{
 *                parent: "fithub.com/",
 *                children: []string{}
 *                depth: 3,
 *            },
 *        },
 *    }
 *
 *    func (f *fakeFetcher) Fetch(url string, resc chan *result, errc chan error, m *mockUrl) {
 *        if (len(url)==0 || url==""){
 *            e := errors.New("No Urls Found")
 *            errc<- e
 *        }
 *        
 *        res := &result{
 *            url: m.parent,
 *            child: getChild(url),
 *            depth: getDepth(url),
 *        }
 *
 *        resc<- res
 *    }
 *}
 */

 //Rendition 1

/*
 *func TestCrawl(t *testing.T){
 *    type fakeFetcher map[string]*fakeResult
 *    type fakeResult struct {
 *        body string
 *        urls []string
 *    }
 *
 *    func (f fakeFetcher) Fetch(url string) (string, []string, error) {
 *        if val, ok := f[url]; ok {
 *            return val.body, val.urls, nil
 *        }
 *        return "", nil, fmt.Errorf("not found: %s", url)
 *    }
 *    var fetcher = fakeFetcher{
 *        "http://golang.org/": &fakeResult{
 *            "Go Website",
 *            []string{
 *                "http://golang.org/pkg/",
 *                "http://golang.org/os/",
 *            },
 *        },
 *        "http://golang.org/pkg/": &fakeResult{
 *            "Packages",
 *            []string{
 *                "http://golang.org/pkg/os/",
 *                "http://golang.org/pkg/fmt/",
 *                "http://golang.org/",
 *            },
 *        },
 *        "http://golang.org/os/": &fakeResult{
 *            "OS stuff",
 *            []string{
 *                "http://golang.org/os/detect",
 *                "http://golang.org/pkg",
 *                "http://golang.org/",
 *            },
 *        },
 *    }
 *
 *    go func(){
 *        
 *    }()
 *}
 */
