#! /bin/bash

identification_regex='^[0-9]{10}$'
countryRegex='EC|CO|US'
dateRegex='^19|20[0-9]{2}[1-12][1-1;31]'

identification=''
country=''
date=''

read -p "What your identification? " identification
read -p "What are you from?, CO, EC, US " country
read -p "What day were you born? yyyyMMDD " date
echo ""

if [[ $identification =~ $identification_regex ]]; then
    echo -e "Identification ($identification) status \e[32m APPROVED \e[0m"
else
    echo -e "Identification ($identification) status \e[31mDENIED\e[0m"
fi

if [[ $country =~ $countryRegex ]]; then
    echo -e "Country ($country) status \e[32m APPROVED \e[0m"
else
    echo -e "Country ($country) status \e[31mDENIED\e[0m"
fi

if  [[ $date =~ $dateRegex ]]; then
    echo -e "Date ($date) status \e[32m APPROVED \e[0m"
else
    echo -e "Date ($date) status \e[31mDENIED\e[0m"
fi

if [[ $identification =~ $identification_regex && $country =~ $countryRegex 
    && $date =~ $dateRegex ]]; then
    echo -e "Your identification is \e[32m $identification \e[0m in the country \e[32m $country \e[0m the date \e[32m $date \e[0m"
fi

