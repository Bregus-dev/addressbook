# addressbook
console utility for working with the address book of iBregus components

```bash
go build -o addressbook.bin
```

```bash
./addressbook.bin
```

```bash
./addressbook.bin --help
```

```bash
./addressbook.bin info
```

---

```bash
./addressbook.bin read --help
```

```bash
./addressbook.bin read --fileFrom bregus_iron_addressbook.yml
```

```bash
./addressbook.bin read --fileFrom bregus_iron_addressbook.yml --mode all
```

```bash
./addressbook.bin read --fileFrom bregus_iron_addressbook.yml --motherboard "motherboard fab.02"
```

```bash
./addressbook.bin read --fileFrom bregus_iron_addressbook.yml --motherboard "motherboard fab.02" --mode all
```

```bash
./addressbook.bin read --fileFrom bregus_iron_addressbook.yml --module "TurbModul fab.03.A"
```

```bash
./addressbook.bin read --fileFrom bregus_iron_addressbook.yml --driver "White modul fab.01"
```

```bash
./addressbook.bin read --fileFrom bregus_iron_addressbook.yml --module "current loop input 1.0" --mode all
```

---

```bash
./addressbook.bin write-go --help
```

```bash
./addressbook.bin write-go --fileFrom bregus_iron_addressbook.yml --fileTo "./temp"
```

---

```bash
./addressbook.bin write-json --help
```

```bash
./addressbook.bin write-json --fileFrom bregus_iron_addressbook.yml --fileTo "./temp"
```

---

**Жир**

```bash
./addressbook.bin write-md --help
```

```bash
./addressbook.bin write-md --fileFrom bregus_iron_addressbook.yml --fileTo "./temp"
```

---

```bash
./addressbook.bin write-cpp --help
```

```bash
./addressbook.bin write-cpp --fileFrom bregus_iron_addressbook.yml --fileTo "./temp"
```

```bash
./addressbook.bin write-cpp --fileFrom bregus_iron_addressbook.yml --fileTo "./temp/filemane.cpp"
```
