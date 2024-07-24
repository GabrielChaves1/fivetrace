import axios from "axios";
import { useEffect, useMemo, useState } from "react";
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from "./ui/select";

type Country = {
  name: {
    common: string;
  }
  flags: {
    png: string;
  }
}

export default function CountrySelector() {
  const countries = [
    {
      name: {
        common: 'Brasil'
      },
      flags: {
        png: 'https://flagcdn.com/br.svg'
      }
    }
  ];

  const sortedCountries = useMemo(() => countries.sort((a, b) => a.name.common.localeCompare(b.name.common)), [countries]);

  return (
    <Select defaultValue="Brasil">
      <SelectTrigger className="w-full">
        <SelectValue placeholder="Selecione um país" />
      </SelectTrigger>
      <SelectContent>
        <SelectItem value="none" disabled>
          Selecione um país
        </SelectItem>
        {sortedCountries.map(country => (
          <SelectItem key={country.name.common} value={country.name.common}>
            <div className="flex gap-2">
              <img src={country.flags.png} alt={country.name.common} className="w-5 h-5 object-contain" />
              {country.name.common}
            </div>
          </SelectItem>
        ))}
      </SelectContent>
    </Select>
  )
}