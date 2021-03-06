# Запуск цієї программи порахуує хеш та надрукує його у
# форматі зрозумілому для людського ока.
$ go run sha1-hashes.go
це рядок sha1
afa57c30c2333a0a4bccd5d7292efc211f07e510
використай sha1.Sum([]byte)
5f0d608c5a8be56c733b6113d6161198f687436c


# Ви можете обчислювати і інші хеші використовуючи схожий
# до вище приведеного сценарій. Наприклад, щоб порахувати
# MD5 хеші імпортуйте `crypto/md5` і скористайтесь `md5.New()`
# (або `md5.Sum([]byte)` що трошка коротше).

# Зауважте що якщо вам потрібні криптографічно безпечні
# хеші вам потрібно провести уважне дослідження
# [їх надійності](https://uk.wikipedia.org/wiki/Криптографічна_хеш-функція)!
