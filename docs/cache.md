
# NeoCheckin - Cache

**A layer between the [wrapper server](wrapper.md) and the frontend. For faster response and removal of dependency on VPN and internet connection for data integrity.**

## Technologies

- [**Go** (Golang)](https://golang.org/)

## Exported Models

### Employee

```ts
{
    name: string,
    flex: number,
    working: boolean,
    department: string,
    photo: string, // Base64 image url
}
```

### Option

Option:
```ts
{
    id: number,
    name: string
}
```

## REST Api

### Errors
In case of a status code `>= 400` this will be the response instead.
```ts
{
    error: string
}
```

### GET `/api/employee/:rfid`
#### Request
```ts
{
    rfid: string
}
```
#### Response
```ts
{
    employee: Employee,
}
```

### POST `/api/employee/cardscanned`
#### Request
```ts
{
    employeeRfid: string,
    checkingIn: boolean
    optionId: number // thies peiter fuckery on this one, please fix
}
```
#### Response
```ts
{
    employee: Employee
}
```

### GET `/api/employees/working`

#### Response
```ts
{
    employees: Employee[],
    ordered: {
        [department: string]: Employee[]
    }
}
```

### GET `/api/options/available`
#### Response
```ts
{
    options: Option[],
}
```


