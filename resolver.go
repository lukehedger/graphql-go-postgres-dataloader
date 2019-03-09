package api

import (
	"fmt"
	"strings"
)

func (r *Resolver) Hello(args struct{ ID string }) (string, error) {
	var name string
	rows, err := r.DB.Query("SELECT name FROM people WHERE id = $1 LIMIT $2", args.ID, 1)
	if err != nil {
		return "", err
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&name)
		if err != nil {
			return "", err
		}
	}
	err = rows.Err()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Hello, %v", name), nil
}

func (r *Resolver) People(args struct{ IDs []string }) ([]*PersonResolver, error) {
	var people []*Person
	rows, err := r.DB.Query("SELECT * FROM people WHERE id IN (" + fmt.Sprintf("'%s'", strings.Join(args.IDs, "','")) + ")")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var person Person
		err = rows.Scan(&person.ID, &person.Name)
		if err != nil {
			return nil, err
		}
		people = append(people, &person)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	resolvers := make([]*PersonResolver, len(people))
	for i, person := range people {
		resolvers[i] = &PersonResolver{person}
	}

	return resolvers, nil
}

func (r *Resolver) PeopleViaLoader(args struct{ IDs []string }) ([]*PersonResolver, error) {
	var people []*Person
	rows, err := r.DB.Query("SELECT * FROM people WHERE id IN (" + fmt.Sprintf("'%s'", strings.Join(args.IDs, "','")) + ")")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var person Person
		err = rows.Scan(&person.ID, &person.Name)
		if err != nil {
			return nil, err
		}
		people = append(people, &person)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	resolvers := make([]*PersonResolver, len(people))
	for i, person := range people {
		resolvers[i] = &PersonResolver{person}
	}

	return resolvers, nil
}

func (r *PersonResolver) ID() string {
	return r.Person.ID
}

func (r *PersonResolver) Name() string {
	return r.Person.Name
}
