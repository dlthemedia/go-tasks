package main

import "fmt"

func main() {

	profile := Profile{"Andrey", 67, false}

	fmt.Printf("%+v\n", profile)

	var cfg AppConfig

	fmt.Printf("%+v\n", cfg)
	fmt.Printf("%+v\n", cfg.Host)
	fmt.Printf("%+v\n", cfg.Port)
	fmt.Printf("%+v\n", cfg.Debug)

	point := Address{"Donbass", "Efremov"}
	user := Employee{"Макан", point}
	fmt.Printf("%+v\n", user)

	shipment := Shipment{
		Package: Package{
			ID:     "PKG-001",
			Weight: 15,
		},
		Destination: Destination{
			City: "Москва",
			Zip:  "101000",
		},
	}

	fmt.Printf("%+v\n", shipment)

	article := Article{
		Title: "Встраивание структур в Go",
		Audit: Audit{
			CreatedAt: "2024-01-01",
			UpdatedAt: "2024-01-02",
		},
	}

	// Вывод полей
	fmt.Println("Title:", article.Title)
	fmt.Println("CreatedAt:", article.CreatedAt)
	fmt.Println("UpdatedAt:", article.UpdatedAt)

	// Создаём клиента
	client := Client{
		ID: "CLIENT-001",
		AddresAddress: Address{
			City:   "Москва",
			Street: "Тверская, 10",
		},
		ContactInfo: ContactInfo{
			Phone: "+7-999-123-45-67",
			Email: "ivan@example.com",
		},
	}

	fmt.Println("ID:", client.ID)
	fmt.Println("City:", client.AddresAddress.City) // доступ через обычное поле
	fmt.Println("Email:", client.Email)

	enrollment := CourseEnrollment{
		EnrollmentID: "ENR-2024-001",
		Status:       StatusActive,
		Grade:        0, // ещё не оценено
		Student: Student{
			ID:   "STU-123",
			Name: "Анна Петрова",
			InfoContactInfo: InfoContactInfo{
				Email: "anna@example.com",
				Phone: "+7-999-888-77-66",
			},
		},
		Course: Course{
			Code:        "GO-101",
			Title:       "Введение в Go",
			CreditHours: 3,
		},
	}

	fmt.Println("=== ОТЧЁТ О ЗАПИСИ НА КУРС ===")
	fmt.Printf("ID записи: %s\n", enrollment.EnrollmentID)
	fmt.Printf("Статус: %s\n", enrollment.Status)
	fmt.Printf("Студент: %s (%s)\n", enrollment.Student.Name, enrollment.Student.Email)
	fmt.Printf("Курс: %s - %s\n", enrollment.Course.Code, enrollment.Course.Title)
	fmt.Printf("Текущая оценка: %d\n", enrollment.Grade)

}
