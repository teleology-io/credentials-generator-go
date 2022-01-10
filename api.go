package generator

import (
	"regexp"
	"strings"
)

type UsernameOptions struct {
	Seed       string
	Lowercase  bool
	Adjectives bool
	Animals    bool
	Countries  bool
	Numbers    int
}

func Username(opts UsernameOptions) string {
	out := []string{}

	if opts.Seed != "" {
		out = append(out, opts.Seed)
	}

	if opts.Adjectives {
		it := random(adjectives())
		out = append(out, it)
	}

	if opts.Animals {
		it := random(animals())
		out = append(out, it)
	}

	if opts.Countries {
		it := random(countries())
		out = append(out, it)
	}

	if opts.Numbers != 0 {
		for i := 0; i < opts.Numbers; i++ {
			it := random(numbers())
			out = append(out, it)
		}
	}

	shuffled := shuffle(out)
	joined := strings.Join(shuffled[:], "")

	reg := regexp.MustCompile(`\s+|\'`)
	result := reg.ReplaceAllString(joined, "")
	if opts.Lowercase {
		return strings.ToLower(result)
	}

	return result
}

type PasswordOptions struct {
	Lowercase int
	Special   int
	Uppercase int
	Numbers   int
}

func Password(opts PasswordOptions) string {
	out := []string{}
	if opts.Numbers != 0 {
		for i := 0; i < opts.Numbers; i++ {
			it := random(numbers())
			out = append(out, it)
		}
	}

	if opts.Lowercase != 0 {
		for i := 0; i < opts.Lowercase; i++ {
			it := random(lowercase())
			out = append(out, it)
		}
	}

	if opts.Uppercase != 0 {
		for i := 0; i < opts.Uppercase; i++ {
			it := random(uppercase())
			out = append(out, it)
		}
	}

	if opts.Special != 0 {
		for i := 0; i < opts.Special; i++ {
			it := random(special())
			out = append(out, it)
		}
	}

	shuffled := shuffle(out)
	return strings.Join(shuffled[:], "")
}
