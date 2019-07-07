package main

import (
	"reflect"
	"testing"
)

func Test_setFirstLink(t *testing.T) {
	type args struct {
		args []string
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "Test1", args: args{args: []string{"http://monzo.com"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setFirstLink(tt.args.args)
		})
	}
}

func Test_visit(t *testing.T) {
	type args struct {
		link    string
		baseURL string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{name: "Test1",
			args: args{link: "http://mock.eduardohitek.com/a.html", baseURL: "http://mock.eduardohitek.com"},
			want: []string{"http://mock.eduardohitek.com/b.html", "http://mock.eduardohitek.com/d.html",
				"http://mock.eduardohitek.com/g.html", "http://mock.eduardohitek.com/folder/h.html"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := visit(tt.args.link, tt.args.baseURL); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("visit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_visitLink(t *testing.T) {
	type args struct {
		link    string
		baseURL string
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "Test1", args: args{link: "http://mock.eduardohitek.com", baseURL: "http://mock.eduardohitek.com"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			visitLink(tt.args.link, tt.args.baseURL)
		})
	}
}

func Test_returnAllLinks(t *testing.T) {
	type args struct {
		url     string
		urlBase string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{name: "Test1",
			args: args{
				url:     "http://mock.eduardohitek.com/a.html",
				urlBase: "http://mock.eduardohitek.com",
			},
			want: []string{"http://mock.eduardohitek.com/b.html", "http://mock.eduardohitek.com/d.html",
				"http://mock.eduardohitek.com/g.html", "http://mock.eduardohitek.com/folder/h.html"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := returnAllLinks(tt.args.url, tt.args.urlBase); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("returnAllLinks() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_formatURL(t *testing.T) {
	type args struct {
		base string
		link string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "Test1",
			args: args{
				base: "http://monzo.com",
				link: "/about",
			},
			want: "http://monzo.com/about"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := formatURL(tt.args.base, tt.args.link); got != tt.want {
				t.Errorf("formatURL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_returnLocalLinks(t *testing.T) {
	type args struct {
		baseURL string
		links   []string
	}
	tests := []struct {
		name           string
		args           args
		wantLocalLinks []string
	}{
		{name: "Test1",
			args: args{
				baseURL: "http://monzo.com",
				links:   []string{"http://monzo.com/about", "http://monzo.com/legal", "http://twitter.com/monzo"},
			},
			wantLocalLinks: []string{"http://monzo.com/about", "http://monzo.com/legal"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotLocalLinks := returnLocalLinks(tt.args.baseURL, tt.args.links); !reflect.DeepEqual(gotLocalLinks, tt.wantLocalLinks) {
				t.Errorf("returnLocalLinks() = %v, want %v", gotLocalLinks, tt.wantLocalLinks)
			}
		})
	}
}

func Test_addLinkToList(t *testing.T) {
	type args struct {
		link       string
		linksFound []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{name: "Test1", args: args{link: "http://monzo.com", linksFound: []string{}}, want: []string{"http://monzo.com"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addLinkToList(tt.args.link, tt.args.linksFound); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addLinkToList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_trimHash(t *testing.T) {
	type args struct {
		l string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "Test1",
			args: args{l: "http://monzo.com/about#information"},
			want: "http://monzo.com/about"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := trimHash(tt.args.l); got != tt.want {
				t.Errorf("trimHash() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_removeLastSlash(t *testing.T) {
	type args struct {
		l string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "Test1",
			args: args{l: "http://monzo.com/about/"},
			want: "http://monzo.com/about"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeLastSlash(tt.args.l); got != tt.want {
				t.Errorf("removeLastSlash() = %v, want %v", got, tt.want)
			}
		})
	}
}
