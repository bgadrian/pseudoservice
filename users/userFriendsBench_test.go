package users

import (
	"testing"

	"github.com/brianvoe/gofakeit"
)

type example struct {
	Age      int8   `fake:"#"`
	Name     string `fake:"{person.first} {person.last}"`
	Company  string `fake:"{company.buzzwords} {company.bs} {company.suffix}"`
	Position string `fake:"{job.descriptor} {job.level} {job.title}"`
	Email    string `fake:"@{person.last}.{internet.domain_suffix}"`
}

//reflection should be slower
func BenchmarkGofakeitStruct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		st := example{}
		gofakeit.Seed(42)
		gofakeit.Struct(&st)
	}
}

func BenchmarkGofakeitDirectCall(b *testing.B) {
	for i := 0; i < b.N; i++ {
		st := example{}
		gofakeit.Seed(42)
		st.Age = gofakeit.Int8()
		st.Name = gofakeit.Name()
		st.Company = gofakeit.BuzzWord() + " " +
			gofakeit.BS() + " " + gofakeit.CompanySuffix()
		st.Position = gofakeit.JobDescriptor() + " " +
			gofakeit.JobLevel() + " " +
			gofakeit.JobTitle()
		st.Email = gofakeit.LastName() + "." + gofakeit.DomainSuffix()
	}
}
