# BudgeTea
## About BudgeTea
BudgeTea is a simple tool to manage your budget. The application should be able to store expenses divided into a monthly spending reports where you can also see every expense and how much it cost.

User is able to set their monthly budget, add expenses, view their spending report in the app where they will be able to also export them as a spreadsheat (.csv format).

When user wants to export their report they will have an option to export:
- Expenses for the current month
- Expenses from a past month
- All expenses

## Functionalitiess
### Add an Expense
User will be able to enter:
- Name of the expense (store, item, etc...)
- Category of the expense
  - Rent
  - Other bills
  - Grocieries
  - Travel
  - Hobbies
  - Miscelanious 
- Price (in a selected denomination)
  - Supported denominations:
    - Swedish Crown (SEK)
    - Euros (EUR)
    - American Dollars (USD)

### See Expenses:
The user will be able to see how much money they spent this month, whether they are under or above their selected budget and how much. 
Here the user can also choose whether they want to download their expense report either for a specific month or all time expense report. Individual expenses 
will be seen in these reports.

### Preferences:
The user will be able to edit their budget and preferred denomination.

## Entity Relationship Diagram
Diagram showing between multiple entities in the database


## Tools:
The backend of BudgeTea is built using Go with Fyne as the front end library. BudgeTea is also using JSON to store data.

## Binaries
Binaries (executable files) for Windows and Linux will be found in the ***binaries*** folder

## Build
