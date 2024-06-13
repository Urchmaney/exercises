import "./styles.css";

export default function App() {
  return (
    <>
      <ParentContainer />
    </>
  );
}

function ProfileDetail() {
  return <div className="profile-detail"></div>;
}

function Check() {
  return <div className="check"> </div>;
}

function ProfileContainer() {
  return (
    <div className="profile-container">
      <div className="profile"></div>
    </div>
  );
}

function InnerContainer() {
  return (
    <div className="inner-container">
      <ProfileContainer />
      <ProfileDetail />
      <Check />
    </div>
  );
}

function ParentContainer() {
  return (
    <div className="parent-container">
      <InnerContainer />
    </div>
  );
}
