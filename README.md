# **VoyageBot**

A bot to get flights and B&B prices at a low cost.

## ***The request***

The URL to make the request is composed by these parameters:

    momondo.it/flight-search/FIRST_CITY-SECOND_CITY/DATE_P/DATE_R/N_adults/N_students/N_children-11-17?fs=cfc=1;bfc=1&sort=bestflight_a

- **FIRST_CITY** and **SECOND_CITY** are respectively the city of departure and the city of arrival (if it's a one way flight the SECOND_CITY can be omitted)

- **DATE_P** and **DATE_R** are respectively the date of departure and the date of arrival (if it's a one way flight the DATA_R can be omitted)

- **N_adults** is the number of adults (if there's only 1 adult this parameter can be omitted)

- **N_students** is the number of students (if there's only 1 adult this parameter can be omitted)

- **N_children-11-17** is the number of students (if there's only 1 adult this parameter can be omitted)

- **cfc** and **bfs** are the number of handheld luggages and the number of luggages in the hold

## ***Organizing the scraper:***

### **User input**

- [ ] TODO: Find a way to craft the URL to make the request to momondo.it
