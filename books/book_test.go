package books_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "playground/books"
)

var _ = Describe("Book", func() {
	var (
		longBook  Book
		shortBook Book
	)

	BeforeEach(func() {
		longBook = Book{
			Title:  "Les Miserables",
			Author: "Victor Hugo",
			Pages:  1488,
		}

		shortBook = Book{
			Title:  "Fox In Socks",
			Author: "Dr. Seuss",
			Pages:  24,
		}
	})

	Describe("Categorizing book length", func() {
		Context("With more than 300 pages", func() {
			It("should be a novel", func() {
				Expect(longBook.CategoryByLength()).To(Equal("NOVEL"))
			})
		})

		Context("With fewer than 300 pages", func() {
			It("should be a short story", func() {
				Expect(shortBook.CategoryByLength()).To(Equal("SHORT STORY"))
			})
		})
	})
})

var _ = Describe("Book", func() {
	var book Book

	BeforeEach(func() {
		book, _ = NewBookFromJSON(`{
            "title":"Les Miserables",
            "author":"Victor Hugo",
            "pages":1488
        }`)
	})

	It("can be loaded from JSON", func() {
		Expect(book.Title).To(Equal("Les Miserables"))
		Expect(book.Author).To(Equal("Victor Hugo"))
		Expect(book.Pages).To(Equal(1488))
	})

	It("can extract the author's last name", func() {
		Expect(book.AuthorLastName()).To(Equal("Hugo"))
	})
})

var _ = Describe("Book", func() {
	var (
		book Book
		err  error
	)

	BeforeEach(func() {
		book, err = NewBookFromJSON(`{
            "title":"Les Miserables",
            "author":"Victor Hugo",
            "pages":1488
        }`)
	})

	Describe("loading from JSON", func() {
		Context("when the JSON parses succesfully", func() {
			It("should populate the fields correctly", func() {
				Expect(book.Title).To(Equal("Les Miserables"))
				Expect(book.Author).To(Equal("Victor Hugo"))
				Expect(book.Pages).To(Equal(1488))
			})

			It("should not error", func() {
				Expect(err).NotTo(HaveOccurred())
			})
		})

		Context("when the JSON fails to parse", func() {
			BeforeEach(func() {
				book, err = NewBookFromJSON(`{
                    "title":"Les Miserables",
                    "author":"Victor Hugo",
                    "pages":1488oops
                }`)
			})

			It("should return the zero-value for the book", func() {
				Expect(book).To(BeZero())
			})

			It("should error", func() {
				Expect(err).To(HaveOccurred())
			})
		})
	})

	Describe("Extracting the author's last name", func() {
		It("should correctly identify and return the last name", func() {
			Expect(book.AuthorLastName()).To(Equal("Hugo"))
		})
	})
})

var _ = Describe("Book", func() {
	var (
		book Book
		err  error
		json string
	)

	BeforeEach(func() {
		json = `{
            "title":"Les Miserables",
            "author":"Victor Hugo",
            "pages":1488
        }`
	})

	JustBeforeEach(func() {
		book, err = NewBookFromJSON(json)
	})

	Describe("loading from JSON", func() {
		Context("when the JSON parses succesfully", func() {
			It("should populate the fields correctly", func() {
				Expect(book.Title).To(Equal("Les Miserables"))
				Expect(book.Author).To(Equal("Victor Hugo"))
				Expect(book.Pages).To(Equal(1488))
			})

			It("should not error", func() {
				Expect(err).NotTo(HaveOccurred())
			})
		})

		Context("when the JSON fails to parse", func() {
			BeforeEach(func() {
				json = `{
                    "title":"Les Miserables",
                    "author":"Victor Hugo",
                    "pages":1488oops
                }`
			})

			It("should return the zero-value for the book", func() {
				Expect(book).To(BeZero())
			})

			It("should error", func() {
				Expect(err).To(HaveOccurred())
			})
		})
	})

	Describe("Extracting the author's last name", func() {
		It("should correctly identify and return the last name", func() {
			Expect(book.AuthorLastName()).To(Equal("Hugo"))
		})
	})
})

var a = Measure("it should do something hard efficiently", func(b Benchmarker) {
	runtime := b.Time("runtime", func() {
		output := SomethingHard()
		Expect(output).To(Equal(17))
	})

	Ω(runtime.Seconds()).Should(BeNumerically("<", 0.2), "SomethingHard() shouldn't take too long.")

	b.RecordValue("disk usage (in MB)", HowMuchDiskSpaceDidYouUse())
}, 10)
