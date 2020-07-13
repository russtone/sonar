package models_test

import (
	"net"
	"testing"

	"github.com/miekg/dns"
	"github.com/stretchr/testify/assert"

	"github.com/bi-zone/sonar/internal/models"
)

func TestDNSRRToString(t *testing.T) {
	tests := []struct {
		rec dns.RR
		res string
	}{
		{
			&dns.A{
				Hdr: dns.RR_Header{
					Name:   "sonar.local",
					Rrtype: dns.TypeA,
					Class:  dns.ClassINET,
					Ttl:    uint32(60),
				},
				A: net.ParseIP("127.0.0.1"),
			},
			"127.0.0.1",
		},
		{
			&dns.AAAA{
				Hdr: dns.RR_Header{
					Name:   "sonar.local",
					Rrtype: dns.TypeAAAA,
					Class:  dns.ClassINET,
					Ttl:    uint32(60),
				},
				AAAA: net.ParseIP("2001:0db8:85a3:0000:0000:8a2e:0370:7334"),
			},
			"2001:db8:85a3::8a2e:370:7334",
		},
		{
			&dns.MX{
				Hdr: dns.RR_Header{
					Name:   "sonar.local",
					Rrtype: dns.TypeMX,
					Class:  dns.ClassINET,
					Ttl:    uint32(60),
				},
				Mx:         "example.com.",
				Preference: 10,
			},
			"10 example.com.",
		},
		{
			&dns.TXT{
				Hdr: dns.RR_Header{
					Name:   "sonar.local",
					Rrtype: dns.TypeTXT,
					Class:  dns.ClassINET,
					Ttl:    uint32(60),
				},
				Txt: []string{"txt"},
			},
			"txt",
		},
		{
			&dns.CNAME{
				Hdr: dns.RR_Header{
					Name:   "sonar.local",
					Rrtype: dns.TypeCNAME,
					Class:  dns.ClassINET,
					Ttl:    uint32(60),
				},
				Target: "example.com.",
			},
			"example.com.",
		},
	}

	for _, tt := range tests {
		name := dns.Type(tt.rec.Header().Rrtype).String()

		t.Run(name, func(st *testing.T) {
			assert.Equal(t, tt.res, models.DNSRRToString(tt.rec))
		})
	}
}

func TestDNSStringToRR(t *testing.T) {
	tests := []struct {
		value string
		typ   string
		res   dns.RR
	}{
		{
			"127.0.0.1",
			"A",
			&dns.A{
				Hdr: dns.RR_Header{
					Name:   "test.sonar.local.",
					Rrtype: dns.TypeA,
					Class:  dns.ClassINET,
					Ttl:    uint32(60),
				},
				A: net.ParseIP("127.0.0.1"),
			},
		},
		{
			"2001:db8:85a3::8a2e:370:7334",
			"AAAA",
			&dns.AAAA{
				Hdr: dns.RR_Header{
					Name:   "test.sonar.local.",
					Rrtype: dns.TypeAAAA,
					Class:  dns.ClassINET,
					Ttl:    uint32(60),
				},
				AAAA: net.ParseIP("2001:0db8:85a3:0000:0000:8a2e:0370:7334"),
			},
		},
		{
			"10 example.com.",
			"MX",
			&dns.MX{
				Hdr: dns.RR_Header{
					Name:   "test.sonar.local.",
					Rrtype: dns.TypeMX,
					Class:  dns.ClassINET,
					Ttl:    uint32(60),
				},
				Mx:         "example.com.",
				Preference: 10,
			},
		},
		{
			"txt",
			"TXT",
			&dns.TXT{
				Hdr: dns.RR_Header{
					Name:   "test.sonar.local.",
					Rrtype: dns.TypeTXT,
					Class:  dns.ClassINET,
					Ttl:    uint32(60),
				},
				Txt: []string{"txt"},
			},
		},
		{
			"example.com.",
			"CNAME",
			&dns.CNAME{
				Hdr: dns.RR_Header{
					Name:   "test.sonar.local.",
					Rrtype: dns.TypeCNAME,
					Class:  dns.ClassINET,
					Ttl:    uint32(60),
				},
				Target: "example.com.",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.value, func(st *testing.T) {
			assert.Equal(t, tt.res, models.DNSStringToRR(tt.value, tt.typ, "test", "sonar.local", 60))
		})
	}
}
