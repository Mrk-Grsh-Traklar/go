package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Макдональдс Заказ")
	w.Resize(fyne.NewSize(400, 600))

	title := widget.NewLabel("Оформление заказа")

	// Названия для каждого типа еды
	foodOptions := []string{"Пицца", "Бургер", "Картошка фри", "Напитки"}

	// Объекты для ввода данных
	nameLabel := widget.NewLabel("Ваше имя:")
	nameEntry := widget.NewEntry()

	// Для выбора еды
	foodCheck := widget.NewCheckGroup(foodOptions, func(s []string) {})

	// Для пиццы: размер и количество
	pizzaSizeLabel := widget.NewLabel("Размер пиццы:")
	pizzaSizeRadio := widget.NewRadioGroup([]string{"30 см", "45 см"}, func(s string) {})
	pizzaQuantityLabel := widget.NewLabel("Количество пицц:")
	pizzaQuantityEntry := widget.NewEntry()

	// Для бургера: размер и количество
	burgerSizeLabel := widget.NewLabel("Размер бургера:")
	burgerSizeRadio := widget.NewRadioGroup([]string{"Стандартный", "Гигантский"}, func(s string) {})
	burgerQuantityLabel := widget.NewLabel("Количество бургеров:")
	burgerQuantityEntry := widget.NewEntry()

	// Для картошки фри: размер и количество
	friesSizeLabel := widget.NewLabel("Размер картошки фри:")
	friesSizeRadio := widget.NewRadioGroup([]string{"Маленькая", "Средняя", "Большая"}, func(s string) {})
	friesQuantityLabel := widget.NewLabel("Количество картошек фри:")
	friesQuantityEntry := widget.NewEntry()

	// Для напитков: размер и количество
	drinkSizeLabel := widget.NewLabel("Размер напитка:")
	drinkSizeRadio := widget.NewRadioGroup([]string{"Маленький", "Средний", "Большой"}, func(s string) {})
	drinkQuantityLabel := widget.NewLabel("Количество напитков:")
	drinkQuantityEntry := widget.NewEntry()

	nagetssizeLabel := widget.NewLabel("размер нагетсов:")
	nagetssizeradio := widget.NewRadioGroup([]string{"Маленький", "Средний", "Большой"}, func(s string) {})
	nagetsQualityLabel := widget.NewLabel("Количество нагетсо:")
	nagetsQualityEntry := widget.NewEntry()

	// Для вывода результата
	result := widget.NewLabel("")

	// Кнопка оформления заказа
	orderButton := widget.NewButton("Оформить заказ", func() {
		username := nameEntry.Text
		selectedFood := foodCheck.Selected

		// Формируем текст для вывода с учетом выбранных блюд
		result.SetText("Имя: " + username + "\n")

		if contains(selectedFood, "Пицца") {
			result.SetText(result.Text + "Пицца: " + pizzaSizeRadio.Selected + ", Количество: " + pizzaQuantityEntry.Text + "\n")
		}
		if contains(selectedFood, "Бургер") {
			result.SetText(result.Text + "Бургер: " + burgerSizeRadio.Selected + ", Количество: " + burgerQuantityEntry.Text + "\n")
		}
		if contains(selectedFood, "Картошка фри") {
			result.SetText(result.Text + "Картошка фри: " + friesSizeRadio.Selected + ", Количество: " + friesQuantityEntry.Text + "\n")
		}
		if contains(selectedFood, "Напитки") {
			result.SetText(result.Text + "Напиток: " + drinkSizeRadio.Selected + ", Количество: " + drinkQuantityEntry.Text + "\n")
		}
		if contains(selectedFood, "Нагетсы") {
			result.SetText(result.Text + "Напиток: " + drinkSizeRadio.Selected + ", Количество: " + drinkQuantityEntry.Text + "\n")
		}
	})

	// Добавляем все элементы на экран в контейнер с прокруткой
	scrollContainer := container.NewScroll(container.NewVBox(
		title,
		nameLabel,
		nameEntry,
		foodCheck,
		pizzaSizeLabel,
		pizzaSizeRadio,
		pizzaQuantityLabel,
		pizzaQuantityEntry,
		burgerSizeLabel,
		burgerSizeRadio,
		burgerQuantityLabel,
		burgerQuantityEntry,
		friesSizeLabel,
		friesSizeRadio,
		friesQuantityLabel,
		friesQuantityEntry,
		drinkSizeLabel,
		drinkSizeRadio,
		drinkQuantityLabel,
		drinkQuantityEntry,
		nagetssizeLabel,
		nagetssizeradio,
		nagetsQualityLabel,
		nagetsQualityEntry,
		orderButton,
		result,
	))

	// Устанавливаем прокручиваемый контейнер в окно
	w.SetContent(scrollContainer)

	w.ShowAndRun()
}

// Функция для проверки наличия элемента в срезе
func contains(slice []string, item string) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}
