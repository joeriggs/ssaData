package benefit

/* Contribution and Benefit Base Table.  Taken from:
 *   https://www.ssa.gov/oact/COLA/cbb.html
 *
 * This table contains the maximum earnings for each year.
 */
var maxWages = map[int]int{
  1937:   3000,
  1938:   3000,
  1939:   3000,
  1940:   3000,
  1941:   3000,
  1942:   3000,
  1943:   3000,
  1944:   3000,
  1945:   3000,
  1946:   3000,
  1947:   3000,
  1948:   3000,
  1949:   3000,
  1950:   3000,
  1951:   3600,
  1952:   3600,
  1953:   3600,
  1954:   3600,
  1955:   4200,
  1956:   4200,
  1957:   4200,
  1958:   4200,
  1959:   4800,
  1960:   4800,
  1961:   4800,
  1962:   4800,
  1963:   4800,
  1964:   4800,
  1965:   4800,
  1966:   6600,
  1967:   6600,
  1968:   7800,
  1969:   7800,
  1970:   7800,
  1971:   7800,
  1972:   9000,
  1973:  10800,
  1974:  13200,
  1975:  14100,
  1976:  15300,
  1977:  16500,
  1978:  17700,
  1979:  22900,
  1980:  25900,
  1981:  29700,
  1982:  32400,
  1983:  35700,
  1984:  37800,
  1985:  39600,
  1986:  42000,
  1987:  43800,
  1988:  45000,
  1989:  48000,
  1990:  51300,
  1991:  53400,
  1992:  55500,
  1993:  57600,
  1994:  60600,
  1995:  61200,
  1996:  62700,
  1997:  65400,
  1998:  68400,
  1999:  72600,
  2000:  76200,
  2001:  80400,
  2002:  84900,
  2003:  87000,
  2004:  87900,
  2005:  90000,
  2006:  94200,
  2007:  97500,
  2008: 102000,
  2009: 106800,
  2010: 106800,
  2011: 106800,
  2012: 110100,
  2013: 113700,
  2014: 117000,
  2015: 118500,
  2016: 118500,
  2017: 127200,
  2018: 128400,
  2019: 132900,
  2020: 137700,
  2021: 142800,
}

var meMostRecentYear int = 0
func mostRecentMaxEarningsYear() int {
  if meMostRecentYear == 0 {
    for key, _ := range maxWages {
      if key > meMostRecentYear {
        meMostRecentYear = key
      }
    }
  }
  return meMostRecentYear
}

func MaxEarnings(year int) int {
  if year > mostRecentMaxEarningsYear() {
    return maxWages[mostRecentMaxEarningsYear()]
  } else {
    return maxWages[year]
  }
}

