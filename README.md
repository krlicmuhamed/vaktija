# Go vaktija.ba
Jednostavan golang program koji uzima podatke sa vaktija.ba

# Primjer korištenja:
```bash
l0oky@laptop:~$ vaktija
> Ponedjeljak, 25. April 2016.

Zora 				- 3:50
Izlazak sunca  		- 5:42
Podne 				- 12:45
Ikindija 			- 16:36
Aksam 				- 19:47
Jacija 				- 21:24
l0oky@laptop:~$
```

# Šta ima?
Ovaj program šalje zahtjev na stranici vaktija.ba i čita HTML kako bi uzeo današnje podatke o vremenima namaza.

# Šta fali?
Vaktija.ba nema API tako da sam se poslužio privremenim rješenjem, parsiranje HTML stranice. Još nisam siguran kako ću izabrati drugi grad prilikom slanja zahtjeva.
  - Nemože se birati lokacija po defaultu je samo Sarajevo
  - Ne izbacuje na izlaz još koliko minuta ima do sljedećeg namaza.
  - Nisam napisao postupak instalacije za Windows.
