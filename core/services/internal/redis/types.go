package redismodule

type RedisConfig struct {
	Host       string
	Port       string
	User       string
	Password   string
	ConnectURI string
}

// Нименование HMap в Redis
type RedisHMapNames string

const (
	ListChatMessages       RedisHMapNames = "zodi_list_chat_messages"       // Список сообщений из чатов
	ListDailyHoroscopes    RedisHMapNames = "zodi_list_daily_horoscopes"    // Список ежедневных гороскопов
	ListCompatibility      RedisHMapNames = "zodi_list_compatibility"       // Список совместимостей
	ListPersonalHoroscopes RedisHMapNames = "zodi_list_personal_horoscopes" // Список персональных гороскопов
	ExchangeRateHMapKey    RedisHMapNames = "zodi_exchange_rate"            // Курс валют
	LifeSphereHMapKey      RedisHMapNames = "zodi_life_sphere"              // Сферы жизни
	LanguagesHMapKey       RedisHMapNames = "zodi_languages"                // Языки
)

type RedisMemoryInfo struct {
	UsedMemory             string // Общее количество байтов, выделенных Redis
	UsedMemoryHuman        string // Человеко-читаемое представление UsedMemory
	UsedMemoryRss          string // Количество памяти, которое Redis занимает в оперативной памяти
	UsedMemoryRssHuman     string // Человеко-читаемое представление UsedMemoryRss
	UsedMemoryPeak         string // Максимальное количество памяти, использованное Redis
	UsedMemoryPeakHuman    string // Человеко-читаемое представление UsedMemoryPeak
	UsedMemoryPeakPerc     string // Процент пикового использования памяти от текущего использования
	UsedMemoryOverhead     string // Общее количество памяти, используемое для внутренних механизмов Redis
	UsedMemoryStartup      string // Количество памяти, использованное Redis при запуске
	UsedMemoryDataset      string // Общее количество памяти, используемое для хранения данных
	UsedMemoryDatasetPerc  string // Процент использования памяти, занимаемый данными
	AllocatorAllocated     string // Общее количество памяти, выделенное аллокатором
	AllocatorActive        string // Общее количество памяти, выделенное и активно используемое аллокатором
	AllocatorResident      string // Общее количество памяти, выделенное аллокатором и занимаемое в оперативной памяти
	TotalSystemMemory      string // Общее количество оперативной памяти в системе
	TotalSystemMemoryHuman string // Человеко-читаемое представление TotalSystemMemory
	UsedMemoryLua          string // Общее количество памяти, используемое движком Lua в Redis
	UsedMemoryLuaHuman     string // Человеко-читаемое представление UsedMemoryLua
	Maxmemory              string // Максимальное количество памяти, которое Redis может использовать
	MaxmemoryHuman         string // Человеко-читаемое представление Maxmemory
	MaxmemoryPolicy        string // Политика, которую Redis будет следовать при достижении Maxmemory
	MemFragmentationRatio  string // Отношение памяти RSS к памяти, выделенной аллокатором
	MemFragmentationBytes  string // Разница между памятью RSS и памятью, выделенной аллокатором
	MemAllocator           string // Аллокатор, используемый Redis (jemalloc или glibc)
}
