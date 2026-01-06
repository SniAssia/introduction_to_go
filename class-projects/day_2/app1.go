
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os" 
)

type User struct {
	Name      string 
    Age       int    
    Salary    int   
    Education string 
}

type Data struct {
    Users []User 
}
func (d Data) average() int {
	sum := 0 
	for _, v := range d.Users { 
		sum+=v.Age
	}
    if len(d.Users) > 0 {
	    avg := sum / len(d.Users)
		return avg 

        
    }
	return 0
}

func (d Data) youngest() []string{
    if len(d.Users) == 0 {
        return nil 
	}
	temp := d.Users[0].Age
	for _,v := range d.Users{
		if v.Age < temp {
			temp = v.Age
		}
	}
	names := []string{}
	for _,v := range d.Users{
		if v.Age == temp {
			names = append(names,v.Name)
		}
	}

	return names
}

func (d Data) oldest_per() []string{
    if len(d.Users) == 0 {
        return nil
    }
	temp := d.Users[0].Age
	for _,v := range d.Users{
		if v.Age > temp {
			temp = v.Age
		}
	}
	names := []string{}
	for _,v := range d.Users{
		if v.Age == temp {
			names = append(names,v.Name)
		}
	}

	return names
}

func (d Data) ave_sal()int {
	sum := 0
	for _,v := range d.Users {
		sum+=v.Salary
	}
    if len(d.Users) > 0 {
	    avg := sum / len(d.Users)
        return avg
    }
	return 0
}

func (d Data) highest_sal() []string{
    if len(d.Users) == 0 {
        return nil
    }
	temp := d.Users[0].Salary
	
	for _,v := range d.Users{
		if v.Salary > temp {
			temp = v.Salary
		}
	}
	names := []string{}
	for _,v := range d.Users{
		if v.Salary == temp {
			names = append(names,v.Name)
		}
	}

	return names
}

func (d Data) lowest_sal() []string{
    if len(d.Users) == 0 {
        return nil
    }
	temp := d.Users[0].Salary
	
	for _,v := range d.Users{
		if v.Salary < temp {
			temp = v.Salary
		}
	}
	names := []string{}
	for _,v := range d.Users{
		if v.Salary == temp {
			names = append(names,v.Name)
		}
	}

	return names
}
func (d Data) sorting() map[string]int {
	map1 := make(map[string]int)
	for _,v := range d.Users {
		map1[v.Education]+=1
	}
	return map1
}

func main() {
	filePath := "people.json" 

    jsonData, err := os.ReadFile(filePath) 
	if err != nil {
		log.Fatalf("Error reading file %s: %v", filePath, err)
	}

	var store Data

    
	err = json.Unmarshal(jsonData, &store)
	if err != nil {
		log.Fatalf("Error unmarshalling JSON data: %v", err)
	}

    fmt.Println("\nEducation counts:", store.sorting())
    fmt.Println("Youngest people:", store.youngest())
	fmt.Println("oldest people : ", store.oldest_per())
	fmt.Println("average salary : ",store.ave_sal())
	fmt.Println("average age : ",store.average())
	fmt.Println("people with highest salary : ",store.highest_sal())
	fmt.Println("people with lowest salary : ",store.lowest_sal())
}
