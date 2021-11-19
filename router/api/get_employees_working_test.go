package api_test

import (
	"neocheckin_cache/router/api"
	em "neocheckin_cache/router/api/models/exported_models"
	"testing"
)

func TestSortWorkingEmployees(t *testing.T) {
	{
		i := map[string][]em.Employee{
			"test": {
				em.Employee{
					Name: "aaaa",
				},
				em.Employee{
					Name: "cccc",
				},
				em.Employee{
					Name: "bbbb",
				},
				em.Employee{
					Name: "aabb",
				},
				em.Employee{
					Name: "bbaa",
				},
			},
		}
		e := map[string][]em.Employee{
			"test": {
				em.Employee{
					Name: "aaaa",
				},
				em.Employee{
					Name: "aabb",
				},
				em.Employee{
					Name: "bbaaa",
				},
				em.Employee{
					Name: "bbbb",
				},
				em.Employee{
					Name: "cccc",
				},
			},
		}
		api.SortWorkingEmployees(i)

		for ix := 0; ix < len(i["test"]); ix++ {
			if i["test"][ix].Name != e["test"][ix].Name {
				t.Error("should be sorted alphabetically")
				break
			}
		}

	}
	{
		i := map[string][]em.Employee{
			"test": {
				em.Employee{
					Name: "Theis Pieter Brun Hollebeek",
				},
				em.Employee{
					Name: "Simon Fromse Jakobsen",
				},
				em.Employee{
					Name: "Mikkel Troels Kongsted",
				},
				em.Employee{
					Name: "Maksim Bech Shvets",
				},
				em.Employee{
					Name: "Ole Soelberg",
				},
				em.Employee{
					Name: "Ole Helledie",
				},
			},
		}
		e := map[string][]em.Employee{
			"test": {
				em.Employee{
					Name: "Maksim Bech Shvets",
				},
				em.Employee{
					Name: "Mikkel Troels Kongsted",
				},
				em.Employee{
					Name: "Ole Helledie",
				},
				em.Employee{
					Name: "Ole Soelberg",
				},
				em.Employee{
					Name: "Simon Fromse Jakobsen",
				},
				em.Employee{
					Name: "Theis Pieter Brun Hollebeek",
				},
			},
		}
		api.SortWorkingEmployees(i)

		for ix := 0; ix < len(i["test"]); ix++ {
			if i["test"][ix].Name != e["test"][ix].Name {
				t.Error("should be sorted alphabetically")
				break
			}
		}

	}
}
