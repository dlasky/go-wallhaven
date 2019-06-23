package wallhaven

const baseURL = "https://wallhaven.cc/api/v1"

func getWithBase(p string) string {
	return baseURL + p
}
