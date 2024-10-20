


import javax.swing.JOptionPane;

public class TheatreApp2{
 public static void main(String args []){
// Declare variables
int age, cost;
String day, message;

// Input
day = JOptionPane.showInputDialog(null, "Please enter the day");

age = Integer.parseInt(JOptionPane.showInputDialog(null, "Please enter your age"));


//process
if(day.equalsIgnoreCase("Monday")|| day.equalsIgnoreCase("Tuesday")||
day.equalsIgnoreCase("Wednesday")||day.equalsIgnoreCase("Thursday"))
if(age < 10){
message= "Ticket is Free";
}
else if(age >= 10 && age <= 65){
  message= "Ticket costs E10";
}

else{
      message= "Ticket is Free";
}


}//if(day
else if(day.equalsIgnoreCase("Friday")|| day.equalsIgnoreCase("Saturday")|| day.equalsIgnoreCase("Sunday"))
{

if (age < 10){message = "ticket cost E20";
}
else if (age >= 10 && age <= 65){
  message = "ticket cost E30";
  }else {
    message = "ticket is E10"
    }
else {
message = "Invalid day entered.";
}


//output
//JOptionPane.showMessageDialog(null, message);

 }//main
}//class