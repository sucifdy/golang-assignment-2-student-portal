package main

import (
	"fmt"
	"strings"
)

var Students = []string{
	"A1234_Aditira_TI",
	"B2131_Dito_TK",
	"A3455_Afis_MI",
}

var StudentStudyPrograms = map[string]string{
	"TI": "Teknik Informatika",
	"TK": "Teknik Komputer",
	"SI": "Sistem Informasi",
	"MI": "Manajemen Informasi",
}

type studentModifier func(string, *string)

func Login(id string, name string) string {
	if id == "" || name == "" {
		return "ID or Name is undefined!"
	}
	for _, student := range Students {
		parts := splitStudentData(student)
		if parts[0] == id && parts[1] == name {
			return fmt.Sprintf("Login berhasil: %s", name)
		}
	}
	return "Login gagal: data mahasiswa tidak ditemukan"
}

func Register(id string, name string, major string) string {
	if id == "" || name == "" || major == "" {
		return "ID, Name or Major is undefined!"
	}
	for _, student := range Students {
		if splitStudentData(student)[0] == id {
			return "Registrasi gagal: id sudah digunakan"
		}
	}
	Students = append(Students, fmt.Sprintf("%s_%s_%s", id, name, major))
	return fmt.Sprintf("Registrasi berhasil: %s (%s)", name, major)
}

func GetStudyProgram(code string) string {
	if program, exists := StudentStudyPrograms[code]; exists {
		return program
	}
	return "Kode program studi tidak ditemukan"
}

func ModifyStudent(programStudi, nama string, fn studentModifier) string {
	for i, student := range Students {
		parts := splitStudentData(student)
		if parts[1] == nama {
			fn(programStudi, &Students[i])
			return "Program studi mahasiswa berhasil diubah."
		}
	}
	return "Mahasiswa tidak ditemukan."
}

func UpdateStudyProgram(programStudi string, student *string) {
	parts := splitStudentData(*student)
	*student = fmt.Sprintf("%s_%s_%s", parts[0], parts[1], programStudi)
}

func splitStudentData(data string) []string {
	return strings.Split(data, "_")
}

func main() {
}
