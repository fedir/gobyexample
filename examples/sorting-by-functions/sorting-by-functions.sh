# Запуск нашої программи покаже список фруктів
# відсортований за довжиною рядка.
$ go run sorting-by-functions.go
[ківі банан персик]

# Створючи новий тип данних, реалізуючи три методи
# інтерфейсу `sort.Interface` по відношенню до цього
# типу данних і викликаючи sort.Sort з новоствореним
# типом данних переданим їй аргументом - ми можемо
# сортувати власні типи використовуючи власні методи.
